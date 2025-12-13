defmodule Day12 do
  @moduledoc """
  Advent of Code - Day 12
  """

  defstruct [
    :shapes,           # %{id => MapSet of coords}
    :transformations,  # %{id => [MapSet, MapSet, ...]} (deduplicated)
    :bounds,          # %{id => {width, height}}
    :solid_counts     # %{id => count of # cells}
  ]

  @doc """
  Solves Part 1
  """
  def part1(input) do
    {solver_data, regions} = parse(input)

    regions
    |> Task.async_stream(
      fn {{width, height}, requirements} ->
        can_fit_all_shapes?(width, height, requirements, solver_data)
      end,
      max_concurrency: System.schedulers_online(),
      timeout: 15_000
    )
    |> Enum.count(fn
      {:ok, true} -> true
      _ -> false
    end)
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    IO.puts(input)
    0
  end

  @doc """
  Parse input and prepare solver data structures
  """
  def parse(input) do
    parts =
      input
      |> String.replace("\r", "")
      |> String.split("\n\n", trim: true)

    shapes = into_shapes(Enum.slice(parts, 0..-2//1))
    regions = into_regions(List.last(parts))

    # Precompute all transformations
    transformations = precompute_transformations(shapes)
    bounds = compute_bounds(transformations)
    solid_counts = compute_solid_counts(shapes)

    solver_data = %__MODULE__{
      shapes: shapes,
      transformations: transformations,
      bounds: bounds,
      solid_counts: solid_counts
    }

    {solver_data, regions}
  end

  @doc """
  Check if all required shapes can fit in the region
  """
  def can_fit_all_shapes?(width, height, requirements, solver_data) do
    # Quick feasibility check first
    if not passes_cell_count_check?(width, height, requirements, solver_data.solid_counts) do
      false
    else
      # Build list of shapes to place
      shapes_to_place = expand_requirements(requirements, solver_data.transformations)

        # Try to place with backtracking
        backtrack(MapSet.new(), shapes_to_place, width, height, solver_data)
    end
  end

  defp passes_cell_count_check?(width, height, requirements, solid_counts) do
    total_cells = width * height

    needed_cells =
      requirements
      |> Enum.map(fn {shape_id, count} -> count * Map.get(solid_counts, shape_id, 0) end)
      |> Enum.sum()

    needed_cells <= total_cells
  end

  defp expand_requirements(requirements, transformations) do
    # Convert requirements map to list of shape variants
    # Filter out shapes with count 0
    requirements
    |> Enum.filter(fn {_id, count} -> count > 0 end)
    |> Enum.flat_map(fn {shape_id, count} ->
      variants = Map.get(transformations, shape_id, [])
      List.duplicate({shape_id, variants}, count)
    end)
    # Sort by number of variants (ascending) - place most constrained first
    |> Enum.sort_by(fn {_id, variants} -> length(variants) end)
  end

  @doc """
  Backtracking solver - try to place all shapes
  """
  def backtrack(_grid, [], _width, _height, _solver_data), do: true

  def backtrack(grid, [{_shape_id, shape_variants} | rest], width, height, solver_data) do
    # Try each transformation of this shape
    Enum.any?(shape_variants, fn coords ->
      # Try positions using precomputed list of coords
      try_positions_for_shape(grid, coords, width, height, rest, solver_data)
    end)
  end

  defp try_positions_for_shape(grid, shape_coords, width, height, rest, solver_data) do
    # Convert to list once for reuse
    coords_list = MapSet.to_list(shape_coords)

    # Get bounds
    {min_x, min_y, max_x, max_y} = get_bounds_from_list(coords_list)
    shape_width = max_x - min_x + 1
    shape_height = max_y - min_y + 1

    # Early exit if shape too large
    if shape_width > width or shape_height > height do
      false
    else
      # Try each valid position
      try_each_position(grid, coords_list, min_x, min_y, width, height, shape_width, shape_height, rest, solver_data)
    end
  end

  defp try_each_position(grid, coords_list, min_x, min_y, width, height, shape_width, shape_height, rest, solver_data) do
    # Use a more efficient approach - try positions in order
    max_y_pos = height - shape_height
    max_x_pos = width - shape_width

    do_try_positions(grid, coords_list, min_x, min_y, 0, 0, max_x_pos, max_y_pos, width, height, rest, solver_data)
  end

  defp do_try_positions(_grid, _coords_list, _min_x, _min_y, _x, y, _max_x, max_y, _width, _height, _rest, _solver_data) when y > max_y do
    false
  end

  defp do_try_positions(grid, coords_list, min_x, min_y, x, y, max_x, max_y, width, height, rest, solver_data) when x > max_x do
    do_try_positions(grid, coords_list, min_x, min_y, 0, y + 1, max_x, max_y, width, height, rest, solver_data)
  end

  defp do_try_positions(grid, coords_list, min_x, min_y, x, y, max_x, max_y, width, height, rest, solver_data) do
    # Calculate offset
    dx = x - min_x
    dy = y - min_y

    # Check if this position works
    if can_place_at?(grid, coords_list, dx, dy, width, height) do
      # Place and continue
      new_grid = place_at(grid, coords_list, dx, dy)

      if backtrack(new_grid, rest, width, height, solver_data) do
        true
      else
        # Try next position
        do_try_positions(grid, coords_list, min_x, min_y, x + 1, y, max_x, max_y, width, height, rest, solver_data)
      end
    else
      # Try next position
      do_try_positions(grid, coords_list, min_x, min_y, x + 1, y, max_x, max_y, width, height, rest, solver_data)
    end
  end

  defp get_bounds_from_list(coords_list) do
    xs = Enum.map(coords_list, fn {x, _y} -> x end)
    ys = Enum.map(coords_list, fn {_x, y} -> y end)
    {Enum.min(xs), Enum.min(ys), Enum.max(xs), Enum.max(ys)}
  end

  defp can_place_at?(grid, coords_list, dx, dy, width, height) do
    Enum.all?(coords_list, fn {x, y} ->
      new_x = x + dx
      new_y = y + dy
      new_x >= 0 and new_x < width and
      new_y >= 0 and new_y < height and
      not MapSet.member?(grid, {new_x, new_y})
    end)
  end

  defp place_at(grid, coords_list, dx, dy) do
    Enum.reduce(coords_list, grid, fn {x, y}, acc ->
      MapSet.put(acc, {x + dx, y + dy})
    end)
  end

  defp into_shapes(input) do
    input
    |> Enum.reduce(%{}, fn x, acc ->
      [hd | rest] = String.split(x, ~r/\r\n|\n|\r/, trim: true)
      key = hd |> String.slice(0..-2//1) |> String.to_integer()
      shape_chars = rest |> Enum.map(&String.to_charlist/1)

      # Convert to MapSet of {x, y} coordinates (only # cells)
      coords = shape_to_coords(shape_chars)

      Map.put(acc, key, coords)
    end)
  end

  defp shape_to_coords(shape_chars) do
    for {row, y} <- Enum.with_index(shape_chars),
        {char, x} <- Enum.with_index(row),
        char == ?# do
      {x, y}
    end
    |> MapSet.new()
  end

  defp into_regions(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.map(fn line ->
      [dims, counts] = String.split(line, ":", trim: true)

      # Parse dimensions
      dimensions =
        dims
        |> String.split("x")
        |> Enum.map(&String.to_integer/1)
        |> List.to_tuple()

      # Parse requirements as map {shape_id => count}
      requirements =
        counts
        |> String.split(" ", trim: true)
        |> Enum.map(&String.to_integer/1)
        |> Enum.with_index()
        |> Enum.into(%{}, fn {count, idx} -> {idx, count} end)

      {dimensions, requirements}
    end)
  end

  @doc """
  Precompute all transformations (rotations + flips) for each shape
  """
  def precompute_transformations(shapes) do
    Map.new(shapes, fn {id, coords} ->
      variants = generate_all_transformations(coords)
      {id, variants}
    end)
  end

  defp generate_all_transformations(coords) do
    [
      coords,                                    # Original
      rotate_90(coords),                         # 90° rotation
      rotate_180(coords),                        # 180° rotation
      rotate_270(coords),                        # 270° rotation
      flip_horizontal(coords),                   # Horizontal flip
      flip_vertical(coords),                     # Vertical flip
      coords |> flip_horizontal() |> rotate_90(), # Flip + 90°
      coords |> flip_horizontal() |> rotate_270() # Flip + 270°
    ]
    |> Enum.map(&normalize/1)
    |> Enum.uniq()
  end

  defp rotate_90(coords) do
    MapSet.new(coords, fn {x, y} -> {-y, x} end)
  end

  defp rotate_180(coords) do
    MapSet.new(coords, fn {x, y} -> {-x, -y} end)
  end

  defp rotate_270(coords) do
    MapSet.new(coords, fn {x, y} -> {y, -x} end)
  end

  defp flip_horizontal(coords) do
    MapSet.new(coords, fn {x, y} -> {-x, y} end)
  end

  defp flip_vertical(coords) do
    MapSet.new(coords, fn {x, y} -> {x, -y} end)
  end

  defp normalize(coords) do
    # Move shape to origin (0, 0)
    if MapSet.size(coords) == 0 do
      coords
    else
      coords_list = MapSet.to_list(coords)
      min_x = Enum.map(coords_list, fn {x, _} -> x end) |> Enum.min()
      min_y = Enum.map(coords_list, fn {_, y} -> y end) |> Enum.min()

      MapSet.new(coords, fn {x, y} -> {x - min_x, y - min_y} end)
    end
  end

  @doc """
  Compute bounding boxes for all shape transformations
  """
  def compute_bounds(transformations) do
    Map.new(transformations, fn {id, variants} ->
      # Use the first variant to get bounds (all have same # of cells)
      first_variant = List.first(variants, MapSet.new())
      bounds = if MapSet.size(first_variant) > 0 do
        coords_list = MapSet.to_list(first_variant)
        {min_x, min_y, max_x, max_y} = get_bounds_from_list(coords_list)
        {max_x - min_x + 1, max_y - min_y + 1}
      else
        {0, 0}
      end
      {id, bounds}
    end)
  end

  @doc """
  Compute solid cell counts for each shape
  """
  def compute_solid_counts(shapes) do
    Map.new(shapes, fn {id, coords} ->
      {id, MapSet.size(coords)}
    end)
  end

  @doc """
  Reads input file and runs both parts
  """
  def solve do
    input = read_input()

    IO.puts("--- Part One ---")
    IO.puts("Result: #{part1(input)}")

    IO.puts("--- Part Two ---")
    IO.puts("Result: #{part2(input)}")
  end

  defp read_input do
    File.read!("inputs/day12.txt")
  end
end

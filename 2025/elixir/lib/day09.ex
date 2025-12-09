defmodule Day09 do
  @moduledoc """
  Advent of Code - Day 09
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    coordinates = into_coordinates(input)
    coordinates
    |> Enum.with_index()
    |> Task.async_stream(
      fn {point, i} ->
        Enum.slice(coordinates, (i+1)..-1//1)
        |> Enum.map(fn other_point ->
          area(point, other_point)
        end)
        |> Enum.max(fn -> 0 end)
    end, max_concurrency: System.schedulers_online()
    )
    |> Enum.map(fn {:ok, max} -> max end)
    |> Enum.max()
  end


  @doc """
  Solves Part 2 using coordinate compression and 2D prefix sums
  """
  def part2(input) do
    coordinates = into_coordinates(input)

    # Extract unique rows and columns, then sort them
    rows = coordinates |> Enum.map(fn {_, y} -> y end) |> Enum.uniq() |> Enum.sort()
    cols = coordinates |> Enum.map(fn {x, _} -> x end) |> Enum.uniq() |> Enum.sort()

    # Create mappings from actual coordinates to compressed indices
    row_to_idx = rows |> Enum.with_index() |> Map.new()
    col_to_idx = cols |> Enum.with_index() |> Map.new()

    # Create compressed grid (all false initially)
    compressed_rows = length(rows)
    compressed_cols = length(cols)
    grid = create_grid(compressed_rows, compressed_cols)

    # Mark the polygon path on the compressed grid
    grid = mark_polygon_path(coordinates, grid, row_to_idx, col_to_idx)

    # Fill interior cells (cells not reachable from exterior)
    grid = fill_interior(grid, compressed_rows, compressed_cols)

    # Build 2D prefix sum array for fast rectangle queries
    prefix_sum = build_prefix_sum(grid, compressed_rows, compressed_cols)

    # Find maximum area rectangle
    max_area = find_max_rectangle(coordinates, rows, cols, row_to_idx, col_to_idx, prefix_sum)

    max_area
  end

  defp create_grid(rows, cols) do
    for r <- 0..(rows - 1), into: %{} do
      {r, for(c <- 0..(cols - 1), into: %{}, do: {c, false})}
    end
  end

  defp mark_polygon_path(coordinates, grid, row_to_idx, col_to_idx) do
    n = length(coordinates)

    Enum.reduce(0..(n - 1), grid, fn i, acc_grid ->
      start_coord = Enum.at(coordinates, i)
      end_coord = Enum.at(coordinates, rem(i + 1, n))
      mark_path(start_coord, end_coord, acc_grid, row_to_idx, col_to_idx)
    end)
  end

  defp mark_path({x1, y1}, {x2, y2}, grid, row_to_idx, col_to_idx) do
    r1 = row_to_idx[y1]
    c1 = col_to_idx[x1]
    r2 = row_to_idx[y2]
    c2 = col_to_idx[x2]

    cond do
      r1 == r2 ->
        # Horizontal path
        min_c = min(c1, c2)
        max_c = max(c1, c2)
        Enum.reduce(min_c..max_c, grid, fn c, g ->
          put_in(g[r1][c], true)
        end)

      c1 == c2 ->
        # Vertical path
        min_r = min(r1, r2)
        max_r = max(r1, r2)
        Enum.reduce(min_r..max_r, grid, fn r, g ->
          put_in(g[r][c1], true)
        end)

      true ->
        grid
    end
  end

  defp fill_interior(grid, rows, cols) do
    # BFS from exterior edges to mark all exterior cells
    initial_queue =
      for r <- 0..(rows - 1), c <- 0..(cols - 1),
          (r == 0 or r == rows - 1 or c == 0 or c == cols - 1) and not grid[r][c],
          do: {r, c}

    visited = bfs_exterior(initial_queue, grid, rows, cols, MapSet.new(initial_queue))

    # Fill non-visited, non-path cells (these are interior)
    for r <- 0..(rows - 1), c <- 0..(cols - 1), reduce: grid do
      acc_grid ->
        if not MapSet.member?(visited, {r, c}) and not grid[r][c] do
          put_in(acc_grid[r][c], true)
        else
          acc_grid
        end
    end
  end

  defp bfs_exterior([], _grid, _rows, _cols, visited), do: visited
  defp bfs_exterior([{r, c} | rest], grid, rows, cols, visited) do
    neighbors = [{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}]

    {new_queue, new_visited} =
      Enum.reduce(neighbors, {rest, visited}, fn {nr, nc}, {queue, vis} ->
        if nr >= 0 and nr < rows and nc >= 0 and nc < cols and
           not MapSet.member?(vis, {nr, nc}) and not grid[nr][nc] do
          {queue ++ [{nr, nc}], MapSet.put(vis, {nr, nc})}
        else
          {queue, vis}
        end
      end)

    bfs_exterior(new_queue, grid, rows, cols, new_visited)
  end

  defp build_prefix_sum(grid, rows, cols) do
    # Build (rows+1) x (cols+1) prefix sum array
    base = for r <- 0..rows, into: %{}, do: {r, for(c <- 0..cols, into: %{}, do: {c, 0})}

    for r <- 1..rows, c <- 1..cols, reduce: base do
      acc ->
        val = if grid[r - 1][c - 1], do: 1, else: 0
        sum = val + acc[r - 1][c] + acc[r][c - 1] - acc[r - 1][c - 1]
        put_in(acc[r][c], sum)
    end
  end

  defp find_max_rectangle(coordinates, rows, cols, row_to_idx, col_to_idx, prefix_sum) do
    sorted_coords = Enum.sort_by(coordinates, fn {x, y} -> {y, x} end)
    n = length(sorted_coords)

    Enum.reduce(0..(n - 2), 0, fn i, max_area ->
      {x1, y1} = Enum.at(sorted_coords, i)
      r1 = row_to_idx[y1] + 1
      c1_idx = col_to_idx[x1] + 1

      Enum.reduce((i + 1)..(n - 1), max_area, fn j, current_max ->
        {x2, y2} = Enum.at(sorted_coords, j)

        # Early termination: theoretical max can't beat current best
        theoretical_max = (abs(y2 - y1) + 1) * (abs(x2 - x1) + 1)
        if theoretical_max <= current_max do
          current_max
        else
          r2 = row_to_idx[y2] + 1
          c2_idx = col_to_idx[x2] + 1

          min_r = min(r1, r2)
          max_r = max(r1, r2)
          min_c = min(c1_idx, c2_idx)
          max_c = max(c1_idx, c2_idx)

          compressed_area = (max_r - min_r + 1) * (max_c - min_c + 1)
          sum = prefix_sum[max_r][max_c] - prefix_sum[min_r - 1][max_c] -
                prefix_sum[max_r][min_c - 1] + prefix_sum[min_r - 1][min_c - 1]

          if sum == compressed_area do
            # All cells in compressed rectangle are valid
            actual_min_row = Enum.at(rows, min_r - 1)
            actual_max_row = Enum.at(rows, max_r - 1)
            actual_min_col = Enum.at(cols, min_c - 1)
            actual_max_col = Enum.at(cols, max_c - 1)

            width = actual_max_col - actual_min_col + 1
            height = actual_max_row - actual_min_row + 1
            area = width * height

            max(area, current_max)
          else
            current_max
          end
        end
      end)
    end)
  end

  def into_coordinates(input) do
    input
    |> String.replace("\r", "")
    |> String.split("\n", trim: true)
    |> Enum.map(fn row ->
      [x, y] =
        row
        |> String.split(",")
        |> Enum.map(&String.to_integer/1)

      {x, y}
    end)
  end

  def area({x1, y1}, {x2, y2}) do
    (abs(x2 - x1) + 1) * (abs(y2 - y1) + 1)
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
    File.read!("inputs/day09.txt")
  end
end

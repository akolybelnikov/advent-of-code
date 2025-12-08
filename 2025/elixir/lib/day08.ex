defmodule Day08 do
  @moduledoc """
  Advent of Code - Day 08
  """
  @doc """
  Solves Part 1
  """
  def part1(input, num \\ 1000) do
    input
    |> into_coordinates()
    |> then(&take_n_closest(&1, num))
    |> BoxGrouper.group_boxes()
    |> Enum.sort(fn c1, c2 -> length(c1) >= length(c2) end)
    |> Enum.take(3)
    |> Enum.product_by(&length/1)
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    points = into_coordinates(input)

    sorted_by_distance(points)
    |> BoxGrouper.find_unifying_connection(points)
    |> then(fn {{box1, box2}, _dist} -> elem(box1, 0) * elem(box2, 0) end)
  end

  def into_coordinates(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.map(fn row ->
      [x, y, z] =
        row
        |> String.split(",")
        |> Enum.map(&String.to_integer/1)

      {x, y, z}
    end)
  end

  def euclidean_distance_3d({x1, y1, z1}, {x2, y2, z2}) do
    dx = x2 - x1
    dy = y2 - y1
    dz = z2 - z1
    :math.sqrt(dx * dx + dy * dy + dz * dz)
  end

  def all_pairs(list) do
    for {a, i} <- Enum.with_index(list),
        b <- Enum.drop(list, i + 1),
        do: {a, b}
  end

  def sorted_by_distance(points) do
    all_pairs(points)
    |> Enum.map(fn {p1, p2} -> {{p1, p2}, euclidean_distance_3d(p1, p2)} end)
    |> Enum.sort(fn {_boxes, dist1}, {_boxes2, dist2} -> dist1 <= dist2 end)
  end

  def take_n_closest(points, num \\ 1000) do
    sorted_by_distance(points)
    |> Enum.map(fn {boxes, _dis} -> boxes end)
    |> Enum.take(num)
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
    File.read!("inputs/day05.txt")
  end
end

defmodule BoxGrouper do
  @type adjacency_map :: %{any() => MapSet.t(any())}
  @type box :: any()

  @spec group_boxes([{box(), box()}]) :: [[box()]]
  def group_boxes(pairs) do
    adjacency = build_adjacency_map(pairs)
    all_boxes = Map.keys(adjacency)

    find_groups(all_boxes, adjacency, MapSet.new(), [])
  end

  @doc """
  Finds the first connection that unifies all boxes into a single circuit.
  Uses Union-Find algorithm to efficiently track connected components.
  Returns the connection pair that completes the circuit.
  """
  @spec find_unifying_connection([{{box(), box()}, float()}], [box()]) ::
          {{box(), box()}, float()} | nil
  def find_unifying_connection(sorted_pairs, all_points) do
    total_boxes = length(all_points)
    # Initialize Union-Find: each box is its own parent, rank 0
    initial_uf = %{
      parent: Map.new(all_points, fn box -> {box, box} end),
      rank: Map.new(all_points, fn box -> {box, 0} end),
      count: total_boxes
    }

    Enum.reduce_while(sorted_pairs, initial_uf, fn {{box1, box2}, _dist} = connection, uf ->
      new_uf = union(uf, box1, box2)

      if new_uf.count == 1 do
        {:halt, connection}
      else
        {:cont, new_uf}
      end
    end)
  end

  # Find with path compression
  defp find(%{parent: parent} = uf, box) do
    case Map.get(parent, box) do
      ^box ->
        {box, uf}

      parent_box ->
        {root, uf} = find(uf, parent_box)
        # Path compression: make box point directly to root
        {root, %{uf | parent: Map.put(uf.parent, box, root)}}
    end
  end

  # Union by rank
  defp union(uf, box1, box2) do
    {root1, uf} = find(uf, box1)
    {root2, uf} = find(uf, box2)

    cond do
      root1 == root2 ->
        # Already in same set
        uf

      Map.get(uf.rank, root1) < Map.get(uf.rank, root2) ->
        # Attach root1 under root2
        %{uf | parent: Map.put(uf.parent, root1, root2), count: uf.count - 1}

      Map.get(uf.rank, root1) > Map.get(uf.rank, root2) ->
        # Attach root2 under root1
        %{uf | parent: Map.put(uf.parent, root2, root1), count: uf.count - 1}

      true ->
        # Equal rank: attach root2 under root1 and increment root1's rank
        %{
          uf
          | parent: Map.put(uf.parent, root2, root1),
            rank: Map.update!(uf.rank, root1, &(&1 + 1)),
            count: uf.count - 1
        }
    end
  end

  @spec build_adjacency_map([{box(), box()}]) :: adjacency_map()
  defp build_adjacency_map(pairs) do
    Enum.reduce(pairs, %{}, fn {a, b}, map ->
      map
      |> Map.update(a, MapSet.new([b]), &MapSet.put(&1, b))
      |> Map.update(b, MapSet.new([a]), &MapSet.put(&1, a))
    end)
  end

  @spec find_groups([box()], adjacency_map(), MapSet.t(box()), [[box()]]) :: [[box()]]
  defp find_groups([], _adjacency, _visited, groups), do: groups

  @dialyzer {:nowarn_function, find_groups: 4}
  defp find_groups([box | rest], adjacency, visited, groups) do
    if MapSet.member?(visited, box) do
      find_groups(rest, adjacency, visited, groups)
    else
      {group, new_visited} = explore_group([box], adjacency, visited, MapSet.new())
      find_groups(rest, adjacency, new_visited, [group | groups])
    end
  end

  @spec explore_group([box()], adjacency_map(), MapSet.t(box()), MapSet.t(box())) ::
          {[box()], MapSet.t(box())}
  defp explore_group([], _adjacency, visited, group), do: {MapSet.to_list(group), visited}

  @dialyzer {:nowarn_function, explore_group: 4}
  defp explore_group([box | rest], adjacency, visited, group) do
    if MapSet.member?(visited, box) do
      explore_group(rest, adjacency, visited, group)
    else
      visited = MapSet.put(visited, box)
      group = MapSet.put(group, box)
      neighbors = Map.get(adjacency, box, MapSet.new()) |> MapSet.to_list()
      explore_group(neighbors ++ rest, adjacency, visited, group)
    end
  end
end

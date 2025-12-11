defmodule Day11 do
  @moduledoc """
  Advent of Code - Day 11
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    outputs = into_outputs(input)
    count_paths(outputs, "you", "out")
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    outputs = into_outputs(input)

    node1 = "dac"
    node2 = "fft"

    path1 =
      count_paths(outputs, "svr", node1) *
      count_paths(outputs, node1, node2) *
      count_paths(outputs, node2, "out")

    path2 =
      count_paths(outputs, "svr", node2) *
      count_paths(outputs, node2, node1) *
      count_paths(outputs, node1, "out")

    path1 + path2
  end

  def count_paths(outputs, start, target) do
    {count, _} = find_path(outputs, %{}, start, target)
    count
  end

  def into_outputs(input) do
    input
    |> String.split(~r/\r\n|\n|\r/, trim: true)
    |> Enum.map(&String.split(&1, ":", trim: true))
    |> Enum.reduce(%{}, fn [key, value], acc ->
      value |> String.split(" ", trim: true) |> then(&Map.put(acc, key, &1))
    end)
  end

  defp find_path(_outputs, paths, current, target) when current == target, do: {1, paths}

  defp find_path(outputs, paths, device, target) do
    if Map.has_key?(paths, device) do
      {Map.get(paths, device), paths}
    else
      neighbors = Map.get(outputs, device, [])

      {count, new_paths} =
        Enum.reduce(neighbors, {0, paths}, fn neighbor, {acc_count, acc_paths} ->
          {n_count, updated_paths} = find_path(outputs, acc_paths, neighbor, target)
          {acc_count + n_count, updated_paths}
        end)

      {count, Map.put(new_paths, device, count)}
    end
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
    File.read!("inputs/day11.txt")
  end
end

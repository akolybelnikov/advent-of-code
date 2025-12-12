defmodule Day12 do
  @moduledoc """
  Advent of Code - Day 12
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    parts =
      input
      |> String.replace("\r", "")
      |> String.split("\n\n", trim: true)

    shapes = into_shapes(Enum.slice(parts, 0..-2//1))
    regions = into_regions(List.last(parts))

    IO.inspect(shapes)
    IO.inspect(regions)

    0
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    IO.puts(input)
    0
  end

  defp into_shapes(input) do
    input
    |> Enum.reduce(%{}, fn x, acc ->
      [hd | rest] = String.split(x, ~r/\r\n|\n|\r/, trim: true)
      key = hd |> String.slice(0..-2//1) |> String.to_integer()
      shape = rest |> Enum.map(&String.to_charlist/1)
      acc |> Map.put(key, shape)
    end)
  end

  defp into_regions(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.reduce([], fn x, acc ->
      [hd, val] = String.split(x, ":", trim: true)

      key =
        hd
        |> String.split("x")
        |> Enum.map(&String.to_integer/1)
        |> Enum.to_list()
        |> List.to_tuple()

      region =
        val
        |> String.split(" ", trim: true)
        |> Enum.map(&String.to_integer/1)
        |> Enum.with_index()

      acc |> List.insert_at(-1, {key, region})
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

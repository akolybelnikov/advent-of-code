defmodule Day04 do
  @moduledoc """
  Advent of Code - Day 04
  """

  defmodule Grid do
    defstruct [:cells, :width, :height]

    def new(input) do
      lines = String.split(input, "\n", trim: true)
      height = length(lines)
      width = lines |> List.first() |> String.length()

      cells =
        lines
        |> Enum.with_index()
        |> Enum.flat_map(fn {line, y} ->
          line
          |> String.to_charlist()
          |> Enum.with_index()
          |> Enum.map(fn {char, x} -> {{x, y}, char} end)
        end)
        |> Map.new()

      %__MODULE__{cells: cells, width: width, height: height}
    end

    def neighbors(%__MODULE__{width: w, height: h}, {x, y}) do
      for dx <- -1..1,
          dy <- -1..1,
          {dx, dy} != {0, 0},
          nx = x + dx,
          ny = y + dy,
          nx >= 0 and nx < w,
          ny >= 0 and ny < h do
        {nx, ny}
      end
    end

    def get(%__MODULE__{cells: cells}, pos), do: Map.get(cells, pos)
  end

  @doc """
  Solves Part 1
  """
  def part1(input) do
    grid = Grid.new(input)

    grid.cells
    |> Enum.count(fn {pos, char} ->
      case char do
        ?@ ->
          count = Grid.neighbors(grid, pos)
          |> Enum.count(fn npos -> Grid.get(grid, npos) == ?@ end)
          count <= 3

        _ ->
          false
      end
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
    File.read!("inputs/day04.txt")
  end
end

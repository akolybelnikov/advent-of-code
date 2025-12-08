defmodule Day06 do
  @moduledoc """
  Advent of Code - Day 06
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    group_by_colums(input)
    |> Enum.reduce(0, fn {_k, [op | nums]}, acc ->
      case op do
        "+" -> acc + Enum.sum(Enum.map(nums, &String.to_integer/1))
        "*" -> acc + Enum.product(Enum.map(nums, &String.to_integer/1))
        _ -> acc
      end
    end)
  end

  defp group_by_colums(input) do
    input
    |> String.split("\n", trim: true)
    |> Stream.map(&String.split/1)
    |> Enum.reduce(%{}, fn row, acc ->
      row
      |> Stream.with_index()
      |> Enum.reduce(acc, fn {val, idx}, map ->
        Map.update(map, idx, [val], &[val | &1])
      end)
    end)
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    input
    |> parse_columns()
    |> Enum.reduce(0, fn {op , nums}, acc ->
      case op do
        "+" -> acc + Enum.sum(nums)
        "*" -> acc + Enum.product(nums)
        _ -> acc
      end
    end)
  end

  def parse_columns(input) do
    lines = String.split(input, "\n", trim: true)
    [ops | number_rows] = Enum.reverse(lines)

    # Step 2: Parse operator row to get column info
    operator_info =
      ops
      |> String.graphemes()
      |> Enum.chunk_by(&(&1 == " "))
      |> Enum.chunk_every(2)
      |> Enum.map(fn
        [[char], spaces] -> {char, length(spaces)}
        [[char]] -> {char, 0}
      end)

    # Step 3: Calculate widths
    widths =
      operator_info
      |> Enum.with_index()
      |> Enum.map(fn {{_op, spaces}, idx} ->
        if idx == length(operator_info) - 1 do
          # Last column: total length - all previous columns
          total_used =
            operator_info
            |> Enum.take(idx)
            # spaces + 1 separator
            |> Enum.map(fn {_, sp} -> sp + 1 end)
            |> Enum.sum()

          String.length(ops) - total_used
        else
          # Other columns use their space count
          spaces
        end
      end)

    # => [3, 3, 3, 3]
    # Calculation: 15 - ((3+1) + (3+1) + (3+1)) = 15 - 12 = 3 ✓

    # Step 4: Calculate starting positions for each column
    positions = [
      0 | Enum.scan(widths |> Enum.take(length(widths) - 1), 0, fn w, acc -> acc + w + 1 end)
    ]

    # => [0, 4, 8, 12]

    # Step 5: Extract columns from number rows
    columns_data =
      number_rows
      |> Enum.map(fn line ->
        positions
        |> Enum.zip(widths)
        |> Enum.map(fn {start, width} ->
          String.slice(line, start, width)
        end)
      end)

    # Step 6: Transpose to group by column
    transposed = columns_data |> Enum.zip() |> Enum.map(&Tuple.to_list/1)

    # Step 7: Extract operators
    operators = Enum.map(operator_info, fn {op, _} -> op end)

    # Step 8: Combine numbers + operator for each column
    Enum.zip(transposed, operators)
    |> Enum.map(fn {nums, op} ->
      ints =
        nums
        |> Enum.map(&String.to_charlist/1)
        |> Enum.reduce(%{}, fn list, acc ->
          list
          |> Enum.with_index()
          |> Enum.reduce(acc, fn {val, idx}, map ->
            Map.update(map, idx, [val], &[val | &1])
          end)
        end)
        |> Map.values()
        |> Enum.map(&List.to_string/1)
        |> Enum.map(&String.trim/1)
        |> Enum.map(&Integer.parse/1)
        |> Enum.map(fn {int, _} -> int end)

      {op, ints}
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
    File.read!("inputs/day05.txt")
  end
end

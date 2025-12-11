defmodule Day10 do
  import Bitwise
  alias Day10.Fraction

  @moduledoc """
  Advent of Code - Day 10
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    input
    |> into_machine()
    |> Enum.map(&solve_lights/1)
    |> Enum.sum()
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    input
    |> into_machine()
    |> Enum.map(&solve_joltage/1)
    |> Enum.sum()
  end

  def into_machine(input) do
    input
    |> String.replace("\r", "")
    |> String.split("\n", trim: true)
    |> Enum.map(fn machine ->
      machine
      |> String.split(" ")
      |> then(fn machine ->
        lights =
          hd(machine)
          |> String.slice(1..-2//1)
          |> String.to_charlist()
          |> Enum.reverse()
          |> Enum.reduce(0, fn char, acc ->
            bit = if char == ?#, do: 1, else: 0
            acc <<< 1 ||| bit
          end)

        buttons =
          Enum.slice(machine, 1..-2//1)
          |> Enum.map(fn button ->
            button
            |> String.slice(1..-2//1)
            |> String.split(",")
            |> Enum.map(&String.to_integer/1)
          end)

        joltage =
          List.last(machine)
          |> String.slice(1..-2//1)
          |> String.split(",")
          |> Enum.map(&String.to_integer/1)
          |> List.to_tuple()

        {lights, buttons, joltage}
      end)
    end)
  end

  def solve_lights({target, buttons, _}) do
    queue = :queue.from_list([{0, 0}])
    visited = MapSet.new([0])

    masks =
      Enum.map(buttons, fn button ->
        Enum.reduce(button, 0, fn idx, mask -> mask ||| 1 <<< idx end)
      end)

    bfs(queue, visited, target, masks)
  end

  defp bfs(queue, visited, target, buttons) do
    case :queue.out(queue) do
      {{:value, {current, steps}}, rest_queue} ->
        if current == target do
          steps
        else
          {next_queue, next_visited} =
            Enum.reduce(buttons, {rest_queue, visited}, fn button, {q, v} ->
              next_state = bxor(current, button)

              if MapSet.member?(v, next_state) do
                {q, v}
              else
                {
                  :queue.in({next_state, steps + 1}, q),
                  MapSet.put(v, next_state)
                }
              end
            end)

          bfs(next_queue, next_visited, target, buttons)
        end

      {:empty, _} ->
        :no_solution
    end
  end

  def solve_joltage({_, buttons, joltage}) do
    num_eqs = tuple_size(joltage)
    num_vars = length(buttons)
    max_val = Tuple.to_list(joltage) |> Enum.max()

    system =
      for i <- 0..(num_eqs - 1) do
        row =
          for btn <- buttons do
            if i in btn, do: Fraction.new(1), else: Fraction.new(0)
          end
        {row, Fraction.new(elem(joltage, i))}
      end

    case solve_linear_system(system, num_vars, max_val) do
      {:ok, solution} -> solution
      :no_solution -> 0
    end
  end

  def solve_linear_system(system, num_vars, max_val) do
    {rref, pivots} = gaussian_elimination(system)

    if inconsistent?(rref) do
      :no_solution
    else
      pivot_indices = Map.keys(pivots) |> MapSet.new()
      free_indices = Enum.to_list(0..(num_vars - 1)) |> Enum.reject(&MapSet.member?(pivot_indices, &1))

      if free_indices == [] do
        sum =
          rref
          |> Enum.map(fn {_, b} ->
            if Fraction.is_integer?(b), do: Fraction.to_integer(b), else: 0
          end)
          |> Enum.sum()
        {:ok, sum}
      else
        find_min_solution(rref, pivots, free_indices, num_vars, max_val)
      end
    end
  end

  def gaussian_elimination(system) do
    num_rows = length(system)
    num_cols = length(elem(hd(system), 0))

    {final_system, pivots, _} =
      Enum.reduce(0..(num_cols - 1), {system, %{}, 0}, fn col_idx, {sys, pivots, current_row} ->
        if current_row >= num_rows do
          {sys, pivots, current_row}
        else
          pivot_row_idx =
            Enum.find_index(Enum.drop(sys, current_row), fn {row, _} ->
              !Fraction.zero?(Enum.at(row, col_idx))
            end)

          case pivot_row_idx do
            nil ->
              {sys, pivots, current_row}

            idx ->
              actual_idx = idx + current_row

              sys = List.replace_at(sys, actual_idx, Enum.at(sys, current_row)) |> List.replace_at(current_row, Enum.at(sys, actual_idx))

              {pivot_row, pivot_b} = Enum.at(sys, current_row)
              pivot_val = Enum.at(pivot_row, col_idx)
              inv_pivot = Fraction.divide(Fraction.new(1), pivot_val)

              norm_row = Enum.map(pivot_row, &Fraction.mul(&1, inv_pivot))
              norm_b = Fraction.mul(pivot_b, inv_pivot)

              sys = List.replace_at(sys, current_row, {norm_row, norm_b})

              sys =
                Enum.with_index(sys)
                |> Enum.map(fn {{row, b}, r_idx} ->
                  if r_idx == current_row do
                    {row, b}
                  else
                    factor = Enum.at(row, col_idx)
                    if Fraction.zero?(factor) do
                      {row, b}
                    else
                      new_row =
                        Enum.zip(row, norm_row)
                        |> Enum.map(fn {v, nv} -> Fraction.sub(v, Fraction.mul(factor, nv)) end)
                      new_b = Fraction.sub(b, Fraction.mul(factor, norm_b))
                      {new_row, new_b}
                    end
                  end
                end)

              {sys, Map.put(pivots, col_idx, current_row), current_row + 1}
          end
        end
      end)

    {final_system, pivots}
  end

  def inconsistent?(rref) do
    Enum.any?(rref, fn {row, b} ->
      Enum.all?(row, &Fraction.zero?/1) and !Fraction.zero?(b)
    end)
  end

  def find_min_solution(rref, pivots, free_indices, _num_vars, max_val) do
    base_cost =
      Enum.reduce(pivots, Fraction.new(0), fn {_, r_idx}, acc ->
        {_, b} = Enum.at(rref, r_idx)
        Fraction.add(acc, b)
      end)

    free_coeffs =
      Enum.map(free_indices, fn k ->
        sum_apk =
          Enum.reduce(pivots, Fraction.new(0), fn {_, r_idx}, acc ->
            {row, _} = Enum.at(rref, r_idx)
            Fraction.add(acc, Enum.at(row, k))
          end)
        {k, Fraction.sub(Fraction.new(1), sum_apk)}
      end)

    pivot_rows =
      Enum.map(pivots, fn {_, r_idx} ->
        {Enum.at(rref, r_idx), 0}
      end)

    case search(free_indices, %{}, pivot_rows, free_coeffs, base_cost, nil, max_val) do
      nil -> :no_solution
      min_cost -> {:ok, Fraction.to_integer(min_cost)}
    end
  end

  def search([], assignment, pivot_rows, free_coeffs, base_cost, current_min, _max_val) do
    valid =
      Enum.all?(pivot_rows, fn {{row, b}, _} ->
        sum_ax =
          Enum.reduce(assignment, Fraction.new(0), fn {k, val}, acc ->
            Fraction.add(acc, Fraction.mul(Enum.at(row, k), Fraction.new(val)))
          end)
        x_p = Fraction.sub(b, sum_ax)

        Fraction.is_integer?(x_p) and Fraction.to_float(x_p) >= -0.0001 and Fraction.to_integer(x_p) >= 0
      end)

    if valid do
      cost =
        Enum.reduce(assignment, base_cost, fn {k, val}, acc ->
          {_, coeff} = Enum.find(free_coeffs, fn {fk, _} -> fk == k end)
          Fraction.add(acc, Fraction.mul(coeff, Fraction.new(val)))
        end)

      if current_min == nil or Fraction.lt?(cost, current_min) do
        cost
      else
        current_min
      end
    else
      current_min
    end
  end

  def search([k | rest], assignment, pivot_rows, free_coeffs, base_cost, current_min, max_val) do
    {min_k, max_k} =
      Enum.reduce(pivot_rows, {0, max_val}, fn {{row, b}, _}, {acc_min, acc_max} ->
        sum_assigned =
          Enum.reduce(assignment, Fraction.new(0), fn {j, val}, acc ->
            Fraction.add(acc, Fraction.mul(Enum.at(row, j), Fraction.new(val)))
          end)
        current_val = Fraction.sub(b, sum_assigned)

        slack =
          Enum.reduce(rest, Fraction.new(0), fn u, acc ->
            a_pu = Enum.at(row, u)
            if Fraction.lt?(a_pu, Fraction.new(0)) do
              Fraction.add(acc, Fraction.mul(a_pu, Fraction.new(max_val)))
            else
              acc
            end
          end)

        limit = Fraction.sub(current_val, slack)
        a_pk = Enum.at(row, k)

        cond do
          Fraction.zero?(a_pk) ->
            if Fraction.lt?(limit, Fraction.new(0)) do
              {1, 0} # Invalid
            else
              {acc_min, acc_max}
            end

          Fraction.lt?(Fraction.new(0), a_pk) ->
            # x_k <= limit / a_pk
            val = Fraction.divide(limit, a_pk)
            new_max = Fraction.floor(val)
            {acc_min, min(acc_max, new_max)}

          true -> # a_pk < 0
            # x_k >= limit / a_pk
            val = Fraction.divide(limit, a_pk)
            new_min = Fraction.ceil(val)
            {max(acc_min, new_min), acc_max}
        end
      end)

    if min_k > max_k do
      current_min
    else
      Enum.reduce(min_k..max_k, current_min, fn val, acc_min ->
        search(rest, Map.put(assignment, k, val), pivot_rows, free_coeffs, base_cost, acc_min, max_val)
      end)
    end
  end

  defmodule Fraction do
    def new(n), do: {n, 1}
    def new(n, d) do
      if d == 0, do: raise "Division by zero"
      g = Integer.gcd(n, d)
      if d < 0, do: {-div(n, g), -div(d, g)}, else: {div(n, g), div(d, g)}
    end
    def add({n1, d1}, {n2, d2}), do: new(n1 * d2 + n2 * d1, d1 * d2)
    def sub({n1, d1}, {n2, d2}), do: new(n1 * d2 - n2 * d1, d1 * d2)
    def mul({n1, d1}, {n2, d2}), do: new(n1 * n2, d1 * d2)
    def divide({n1, d1}, {n2, d2}), do: new(n1 * d2, d1 * n2)
    def to_float({n, d}), do: n / d
    def is_integer?({n, d}), do: d == 1 or rem(n, d) == 0
    def to_integer({n, d}), do: div(n, d)
    def zero?({n, _}), do: n == 0
    def floor({n, d}) do
      if n >= 0 do
        div(n, d)
      else
        if rem(n, d) == 0, do: div(n, d), else: div(n, d) - 1
      end
    end
    def ceil({n, d}) do
      if n >= 0 do
        if rem(n, d) == 0, do: div(n, d), else: div(n, d) + 1
      else
        div(n, d)
      end
    end
    def lt?({n1, d1}, {n2, d2}), do: n1 * d2 < n2 * d1
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
    File.read!("inputs/day10.txt")
  end
end

defmodule LeetCode do
#217.Contains duplicate
  def duplicate?([head | tail]), do: duplicate?(tail, head)
  def duplicate?([], _n), do: :false
  def duplicate?([head | tail], h) do
    if h == head do
      :true
    else
      duplicate?(tail, head)
    end
  end

#217.Contains duplicate(with Enum)
  def e_duplicate?(list) do
    list
    |> Enum.reduce_while(MapSet.new, fn x, c ->
    if MapSet.member?(c, x) do
      {:halt, true}
    else
      {:cont, MapSet.put(c, x)}
    end
    end) == true
  end    
end
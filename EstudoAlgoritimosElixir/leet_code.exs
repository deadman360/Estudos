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
#242. Valid Anagram

  def anagram?(s,t) do
    [s, t]
    |>Enum.map(fn x -> Enum.sort(String.graphemes(x)) end)
    |>Enum.uniq()
    |>then(fn x -> case length(x) do
    1 -> true
    _ -> false
    end
    end)  
  end
#242. Valid Anagram(clean)
  def is_anagram?(s,t), do: sort_charlist(s) == sort_charlist(t)
  defp sort_charlist(x) do
    x
    |>String.replace(" ","")
    |>String.graphemes
    |>Enum.sort
  end
#001.TwoSum
  def two_sum(list, target) do
    Enum.reduce_while(list, MapSet.new(), fn value, cnt ->
    target - value
    |>then(fn x -> 
    if MapSet.member?(cnt, x) do
      {:halt, [x, value]}
    else
      {:cont, MapSet.put(cnt, value)}
    end
    end)
    end)
  end
#001.TwoSum(with tail recursion)
  def r_sum(list, target) do
  r_sum(list, target, MapSet.new())
  end
  defp r_sum([], _target, _mapset), do: :nil
  defp r_sum([head | tail], target, mapset) do
  diff = target - head
    if MapSet.member?(mapset, diff) do
    [diff, head]
    else
    r_sum(tail, target,MapSet.put(mapset, head))
    end
  end
#049.Group arrays
  #Solução Pessima.
  #Transforma lista de string em mapa para que eu possa manipular a charlist livremente
  def transforma_mapa(list1), do: transforma_mapa(list1, [], Map.new())
  def transforma_mapa([], _list2, mapa), do: mapa
  def transforma_mapa(list1, [], mapa), do: transforma_mapa(list1 ,Enum.map(list1, fn x -> Enum.sort(String.to_charlist(x)) end), mapa)
  def transforma_mapa([h1 | t1], [h2 | t2], mapa) do
    x = Map.put(mapa, String.to_atom(h1), h2)
    transforma_mapa(t1, t2, x)
  end
  # Função que separa e agrupa
  def group_array(lista), do: group_array(transforma_mapa(lista),Map.values(transforma_mapa(lista)), [])
  def group_array(_mapa,[],lista), do: lista
  def group_array(mapa,lista_temp, nova_lista) do
    [head | tail] = Enum.uniq(lista_temp)
    dups = Map.filter(mapa, fn {_key, value} -> value == head end)
    lista_dups = Enum.map(Map.keys(dups), fn y -> Atom.to_string(y) end)
    group_array(mapa, tail, [lista_dups | nova_lista])
  end
    
end
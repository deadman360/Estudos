defmodule Gpt do
  def separa_cq(lista) do
    require Integer
    lista |> Enum.filter(&Integer.is_even(&1)) |> Enum.map(&(&1 * &1)) |>
    Enum.sum()
    
  end
  def inverte(string) do
    string
    |> String.upcase()
    |> String.trim()
    |> String.reverse()
  end
  def maps(%{idade: idade}) do
    idade
    |>then(fn x -> cond do
    x >= 18 -> "Maior"
    x < 18 -> "Menor"
    end
    end)
  end
  def soma_preco(lista) do
    lista
    |> Enum.map(fn %{preco: preco} -> preco * 0.9 end)
    |> Enum.sum()
  end
  def fluxo(x, y) do
    cond do
    x && y >= 0 -> x + y
    x || y < 0 -> "Numeros invalidos"
    end
  end
  def filtra(lista) do
    lista
    |>Stream.filter(&(rem(&1,2) !=0))
    |>Stream.map(&(&1 * 2))
    |>Enum.take(5)
  end
  def infinita(n) do
    n
    |>Stream.iterate(&(&1 + 1))
    |>Enum.take(10)
  end
  def fatorial(n), do: fatorial(1,n)
  def fatorial(c, 1), do: c
  def fatorial(c,n) do
    fatorial(c*n , n-1)
  end
 

IO.puts(Gpt.separa_cq([1,2,3,4,5,6]))
IO.inspect(Gpt.inverte("OlAmUNdO"))
IO.inspect(Gpt.maps(%{nome: "jao",idade: 18, email: "jao@jaoajoacom"}))
IO.inspect(Gpt.soma_preco([%{preco: 100}, %{preco: 200}, %{preco: 50}]))
IO.inspect(Gpt.fluxo(5,-3))
IO.inspect(Gpt.filtra([1,2,3,4,5,6,7,8,9,10,11]))
IO.inspect(Gpt.infinita(3))
IO.inspect(Gpt.fatorial(10))
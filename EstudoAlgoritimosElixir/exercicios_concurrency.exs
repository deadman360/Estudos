defmodule Concor do


@moduledoc """
Módulo para estudo de concorrencia
"""

  @doc """
    pid que ouve
  """
  
  def ouve do
    receive do
      {:ok, message} -> IO.puts(message)
    end
  ouve()
  end
  
  @doc """
  função que fala eai, beleza ? 
  a cada 1 segundo
  """
  
  def fala(pid) do
    send(pid, {:ok, "Eai, beleza?"})
    :timer.sleep(1000)
    fala(pid)
  end
  
  @doc """
  função de fatorial para rodar com Task
  """
  
  def fatorial_as(x), do: fatorial_as(x, x)
  def fatorial_as(x, 1), do: x
  def fatorial_as(x, c) do
    fatorial_as(x*(c-1) , c-1)
  end

  @doc """
  função para testar a opção after
  na função receive

  recebe um parametro, numero em
  segundos para o timeout.
  """
  
  def listen_time(x) do
    receive do
    {:ok, message} -> IO.puts("Mensagem recebida")
    after
      x*1000 -> IO.puts("Timeout")
    end
  end
end

defmodule LearnG do
  use GenServer
  require Integer
  
@moduledoc """
  Aprendendo a usar GenServer
  """

  def start_link() do
    GenServer.start_link(__MODULE__, [])
  end

 
  #Client 
  
  def sum(pid) do
    GenServer.call(pid, :sum)
  end

  def add(pid, numero) do
    GenServer.cast(pid, {:add, numero})
  end

  def view(pid) do
    GenServer.call(pid, :view)
  end
  #server
  def init(lista) do
    {:ok, lista}
  end

  def handle_call(:sum, __form, lista) do
    {inteiros, _} = Keyword.pop_values(lista, :numero)
    soma = Enum.sum(inteiros)
  
    {:reply, soma,lista}
       
  end

  def handle_cast({:add, numero}, lista) do
  
    lista_nova = [{:numero, numero}|lista]
    IO.inspect(lista_nova)
    {:noreply, lista_nova}
  end

  def handle_call(:view, __from, lista) do
    {:reply, lista, lista}
  end
end

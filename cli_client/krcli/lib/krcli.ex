defmodule Krcli do
  use Application
  
  # See http://elixir-lang.org/docs/stable/elixir/Application.html
  # for more information on OTP Applications
  def start(_type, _args) do
    import Supervisor.Spec, warn: false

    children = [
      # Define workers and child supervisors to be supervised
      # worker(Krcli.Worker, [arg1, arg2, arg3]),
    ]

    # See http://elixir-lang.org/docs/stable/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Krcli.Supervisor]
    Supervisor.start_link(children, opts)
  end

  def main(args) do
    IO.puts "Running Komorebi CLI client."
    KrOpts.dispatch(args)
  end
end

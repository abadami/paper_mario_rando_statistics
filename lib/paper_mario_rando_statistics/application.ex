defmodule PaperMarioRandoStatistics.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      PaperMarioRandoStatisticsWeb.Telemetry,
      {DNSCluster, query: Application.get_env(:paper_mario_rando_statistics, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: PaperMarioRandoStatistics.PubSub},
      # Start the Finch HTTP client for sending emails
      {Finch, name: PaperMarioRandoStatistics.Finch},
      # Start a worker by calling: PaperMarioRandoStatistics.Worker.start_link(arg)
      # {PaperMarioRandoStatistics.Worker, arg},
      # Start to serve requests, typically the last entry
      PaperMarioRandoStatisticsWeb.Endpoint
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: PaperMarioRandoStatistics.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    PaperMarioRandoStatisticsWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end

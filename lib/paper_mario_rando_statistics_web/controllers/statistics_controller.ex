defmodule PaperMarioRandoStatisticsWeb.StatisticsController do
  use PaperMarioRandoStatisticsWeb, :controller

  def index(conn, _params) do
    render(conn, :index)
  end
end

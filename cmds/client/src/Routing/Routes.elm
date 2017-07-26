module Routing.Routes exposing (Route(..))

type alias PostId =
  Int


type Route
  = HomeRoute
  | PostRoute PostId
  | NotFound

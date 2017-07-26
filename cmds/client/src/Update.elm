module Update exposing (..)

import Navigation
import Models exposing (State)
import Messages exposing (Msg(..))
import Routing.Parsers exposing (urlParser)
import Routing.Routes exposing (..)


update : Msg -> State -> ( State, Cmd Msg )
update msg state =
    case msg of
        ShowHome ->
            ( state, Navigation.newUrl "/" )

        ShowPost postId ->
            ( state, Navigation.newUrl ("/post/" ++ toString postId) )


urlUpdate : Route -> State -> ( State, Cmd Msg )
urlUpdate route state =
    ( { state | route = route }, Cmd.none )

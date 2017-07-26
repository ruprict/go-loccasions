module Routing.Parsers exposing (urlParser, parse)

import Navigation
import String
import UrlParser exposing (Parser, (</>), map, int, oneOf, s, string, parseHash)
import Routing.Routes exposing (..)



parse : Navigation.Location -> Route
parse { pathname } =
    let
        path =
            if String.startsWith "/" pathname then
                String.dropLeft 1 pathname
            else
                pathname
    in
        case UrlParser.parsePath identity routeParser path of
            Err err ->
                NotFound

            Ok route ->
                route


postParser : Parser (Int -> a) a
postParser =
    s "post" </> int


homeParser : Parser a a
homeParser =
    oneOf
        [ (s "elm.html")
        , (s "")
        ]


routeParser : Parser (Route -> a) a
routeParser =
    oneOf
        [ map PostRoute postParser
        , map HomeRoute homeParser
        ]

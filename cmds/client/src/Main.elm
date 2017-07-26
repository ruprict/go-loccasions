module Main exposing (..)

import Html exposing (Html, Attribute, h1, header, a, div, text, button, ul, li, code)
import Html.Attributes exposing (class, href)
import Html.Events exposing (Options, onWithOptions, onClick)
import Navigation exposing (Location)
import UrlParser exposing (..)
import Task


-- MODEL


type alias Model =
    { history : List (Maybe Route)
    , route : Route
    , user : User
    , events : List Event
    }


type alias User =
    { id : String
    , email : String
    , name : String
    }


type alias EventId =
    String


type alias Event =
    { id : EventId
    , name : String
    }


type Route
    = HomeRoute
    | EventRoute EventId
    | EventsRoute
    | NotFoundRoute


initModel : Location -> Model
initModel location =
    Model ([ parsePath matchers location ])
        HomeRoute
        (User "" "" "Anonymous")
        [ Event "" "" ]


init : Location -> ( Model, Cmd Msg )
init location =
    update (NewUrl location.pathname) (initModel location)


main : Program Never Model Msg
main =
    Navigation.program UrlChange
        { init = init
        , view = view
        , update = update
        , subscriptions = (\_ -> Sub.none)
        }



-- ROUTING


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map HomeRoute top
        , map EventRoute (s "events" </> string)
        , map EventsRoute (s "events")
        ]



-- UPDATE


type Msg
    = UrlChange Location
    | NewUrl String


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        UrlChange location ->
            Debug.log ((toString msg))
                ( { model | history = parsePath matchers location :: model.history }
                , Cmd.none
                )

        NewUrl url ->
            ( model
            , Navigation.newUrl url
            )



-- VIEW


viewRoute : Maybe Route -> Html Msg
viewRoute maybeRoute =
    case maybeRoute of
        Nothing ->
            li [] [ text "Invalid URL" ]

        Just route ->
            li [] [ code [] [ text (routeToString route) ] ]


viewLink : String -> Html Msg
viewLink url =
    li [] [ button [ onClick (NewUrl url) ] [ text url ] ]


view : Model -> Html Msg
view model =
    div []
        [ h1 [] [ text "Loccasions" ]
        , ul [] (List.map viewLink [ "/", "/events/", "/events/42" ])
        , ul [] (List.map viewRoute model.history)
        ]


routeToString : Route -> String
routeToString route =
    case route of
        HomeRoute ->
            "home"

        EventRoute id ->
            "Event " ++ (toString id)

        EventsRoute ->
            "list of events"

        NotFoundRoute ->
            "WTF?"

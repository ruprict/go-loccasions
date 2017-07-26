module Tests exposing (..)

import Test exposing (..)
import Expect
import Fuzz exposing (list, int, tuple, string)
import String


all : Test
all =
  describe "Tests"
  [ test "Tests" <|
    \() ->
      Expect.equal "test" "test"
  ]


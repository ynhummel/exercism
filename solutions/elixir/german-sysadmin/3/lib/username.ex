defmodule Username do
  def sanitize([]), do: []
  def sanitize([h | t]), do: normalize_chars(h) ++ sanitize(t)

  defp normalize_chars(char) do
    case char do
      ?ä -> ~c"ae"
      ?ö -> ~c"oe"
      ?ü -> ~c"ue"
      ?ß -> ~c"ss"
      ?_ -> ~c"_"
      letter when letter >= ?a and letter <= ?z -> [letter]
      _ -> []
    end
  end
end

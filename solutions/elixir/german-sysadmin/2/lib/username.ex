defmodule Username do
  def sanitize([]), do: []
  def sanitize([h | t]), do: normalize_chars(h) ++ sanitize(t)

  defp normalize_chars(char) do
    case char do
      ?ä -> [?a, ?e]
      ?ö -> [?o, ?e]
      ?ü -> [?u, ?e]
      ?ß -> [?s, ?s]
      letter when letter >= ?a and letter <= ?z -> [letter]
      letter when letter == ?_ -> [letter]
      _ -> []
    end
  end
end

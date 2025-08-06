defmodule LogParser do
  def valid_line?(line) do
    line =~ ~r/^(\[DEBUG\]|\[INFO\]|\[WARNING\]|\[ERROR\])/
  end

  def split_line(line) do
    Regex.split(~r/<[*-=~]*>/, line)
  end

  def remove_artifacts(line) do
    Regex.replace(~r/end-of-line\d+/i, line, "")
  end

  def tag_with_user_name(line) do
    case Regex.run(~r/User\s*(\S+)/u, line, capture: :all_but_first) do
      [user] -> "[USER] #{user} #{line}"
      nil -> line
    end
  end
end

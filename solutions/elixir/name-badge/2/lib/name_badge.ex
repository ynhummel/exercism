defmodule NameBadge do
  def print(id, name, department) do
    badge_id = if id, do: "[#{id}] - ", else: ""

    badge_department = if department, do: " - #{String.upcase(department)}", else: " - OWNER"

    badge_id <> name <> badge_department
  end
end

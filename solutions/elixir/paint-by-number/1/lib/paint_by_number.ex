defmodule PaintByNumber do
  def palette_bit_size(color_count), do: recursive_size(color_count, 1)

  def recursive_size(color_count, size) do
    if 2 ** size >= color_count do
      size
    else
      recursive_size(color_count, size + 1)
    end
  end

  def(empty_picture()) do
    <<>>
  end

  def test_picture() do
    <<0::2, 1::2, 2::2, 3::2>>
  end

  def prepend_pixel(picture, color_count, pixel_color_index) do
    palet_size = palette_bit_size(color_count)
    <<pixel_color_index::size(palet_size), picture::bitstring>>
  end

  def get_first_pixel(<<>>, _color_count), do: nil

  def get_first_pixel(picture, color_count) do
    palet_size = palette_bit_size(color_count)
    <<value::size(palet_size), _rest::bitstring>> = picture
    value
  end

  def drop_first_pixel(<<>>, _color_count), do: <<>>

  def drop_first_pixel(picture, color_count) do
    palet_size = palette_bit_size(color_count)
    <<_value::size(palet_size), rest::bitstring>> = picture
    rest
  end

  def concat_pictures(picture1, picture2) do
    <<picture1::bitstring, picture2::bitstring>>
  end
end

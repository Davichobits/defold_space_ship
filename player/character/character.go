components {
  id: "character"
  component: "/player/character/character.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"run\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/player/character/character.tilesource\"\n"
  "}\n"
  ""
}

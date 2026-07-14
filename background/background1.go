components {
  id: "background"
  component: "/background/background.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/background/background.atlas\"\n"
  "}\n"
  ""
}

components {
  id: "star"
  component: "/star/star.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"star\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/star/stars.atlas\"\n"
  "}\n"
  ""
}

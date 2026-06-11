components {
  id: "enemy1"
  component: "/enemies/enemy1/enemy1.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"7\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/enemies/enemy1/enemy1.atlas\"\n"
  "}\n"
  ""
}

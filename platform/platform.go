components {
  id: "industrial_platforms"
  component: "/platform/industrial_platforms.tilemap"
  rotation {
    y: 0.020000042
    w: 0.99979997
  }
}
components {
  id: "collisions"
  component: "/platform/collisions.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"/platform/industrial_platforms.tilemap\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"\"\n"
  "mask: \"enemy\"\n"
  ""
}

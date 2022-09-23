components {
  id: "player_control"
  component: "/main/player/player_control.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "player_weapons"
  component: "/main/player/player_weapons.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "rivemodel"
  type: "rivemodel"
  data: "scene: \"/main/player/player_character.rivescene\"\n"
  "default_animation: \"\"\n"
  "material: \"/defold-rive/assets/rivemodel.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "default_state_machine: \"\"\n"
  ""
  position {
    x: 44.0
    y: 155.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "pew"
  type: "factory"
  data: "prototype: \"/main/player/pew.go\"\n"
  "load_dynamically: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "pewpew_shoot"
  type: "sound"
  data: "sound: \"/feedbacks/sounds/pew_shoot.wav\"\n"
  "looping: 0\n"
  "group: \"SFX\"\n"
  "gain: 1.0\n"
  "pan: 0.0\n"
  "speed: 1.0\n"
  "loopcount: 0\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "rocket"
  type: "factory"
  data: "prototype: \"/main/player/rocket.go\"\n"
  "load_dynamically: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "rocket_shoot"
  type: "sound"
  data: "sound: \"/feedbacks/sounds/rocket_shoot.wav\"\n"
  "looping: 0\n"
  "group: \"SFX\"\n"
  "gain: 1.0\n"
  "pan: 0.0\n"
  "speed: 1.0\n"
  "loopcount: 0\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}

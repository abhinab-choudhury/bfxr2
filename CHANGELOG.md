# Changelog

## [Unreleased]

### Bug Fixes
- **DSP sample loop** — Replaced premature `return true` with `break` to prevent empty buffer output.
- **Knob reference error** — Moved event listeners to the correct scope so `parameter_name`, `knob_l`, `knob_r` are defined.
- **Footsteppr global** — Changed `step_terrain = 0` to `this.params.terrain = 0` to avoid creating an implicit global.
- **Keyboard shortcut Ctrl+O** — Corrected handler name from `load_data_button_clicked` to `open_data_button_clicked`.
- **Pitch jump typo** — Fixed `pitch_jump_Speed` → `pitch_jump_repeat_speed` in pickup_coin and explosion templates.
- **onended `this` binding** — Captured `var self = this` to correctly reference the `RealizedSound` instance, not the `AudioBufferSourceNode`.
- **Misc UI** — Fixed garbled tooltips on Apply/Revert Sfx buttons, "quare" → "Square" typo, trailing whitespace in `dropEffect`, missing `check_locked=true` in `generate_hit_hurt`, removed unused `generated_files` array.

### Features
- **Keyboard shortcuts R/M** — Press R to randomize parameters, M to mutate parameters.
- **Transfxr tab** — Enabled previously disabled Transfxr sound generator with a working frequency-sweep square wave synthesizer.
- **Undo/Redo** — Full undo/redo stack (50 levels) for parameter changes; Ctrl+Z (undo), Ctrl+Shift+Z/Ctrl+Y (redo); Undo/Redo buttons in the right panel.
- **Drag-reorder files** — Reorder the file list by dragging items; visual drag-over indicator.

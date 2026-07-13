local M = {}

--- Plays an animation only if it isn't already playing.
--- Prevents restarting the same animation every frame.
--- @param self table Object that stores the current animation state.
--- @param animation hash|string Animation id or name.
function M.play_animation(self, animation)
	if self.current_animation ~= animation then
		self.current_animation = animation
		sprite.play_flipbook("#sprite", animation)
	end
end

return M
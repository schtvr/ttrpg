import { defineStore } from 'pinia'

export const useCharacterStore = defineStore('character', {
  // State is where you define your data
  state: () => ({
    character: {
      name: '',
      skills: [],
      inventory: [],
    },
  }),

  // Actions are for modifying state
  actions: {
    setCharacter(data) {
      this.character = data
    },
    addToInventory(item) {
      this.character.inventory.push(item)
    },
    updateSkill(skill) {
      const index = this.character.skills.findIndex((s) => s.name === skill.name)
      if (index !== -1) {
        this.character.skills[index] = skill
      } else {
        this.character.skills.push(skill)
      }
    },
  },

  // Getters are computed properties for state
  getters: {
    inventoryCount: (state) => state.character.inventory.length,
    characterName: (state) => state.character.name,
  },
})

import { defineStore } from 'pinia'

export const useCharacterStore = defineStore('character', {
  // State is where you define your data
  state: () => ({
    name: 'monkey d. luffy',
    epithet: 'the man who will become king of the pirates',
    skills: {
      agility: 4,
      brawl: 3,
      contacts: 3,
      insight: 2,
      operate: 2,
      perform: 2,
      persuade: 1,
      physique: 1,
      stealth: 1,
      shoot: 1,
    },
    about: 'lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    channels: [{}],
    specials: [{}],
    aspects: [
      {
        name: 'rubber arms',
        description: 'strecthy boy is stretchy',
        effects: [
          {
            skills: ['agility'],
            actions: ['attack', 'defend'],
            modifier: 1,
            channel: 'elemental',
            description: 'lorem ipsum dolor sit amet, consectetur adipiscing elit.',
          },
          {
            skills: ['brawl'],
            actions: ['attack'],
            modifier: 2,
            channel: 'organic',
            description: 'lorem ipsum dolor sit amet, consectetur adipiscing elit.',
          },
        ],
      },
      {
        name: 'walking concussion',
        description: 'meat brain',
        effects: [
          {
            skills: ['stealth', 'insight'],
            actions: ['empower'],
            modifier: -1,
            channel: '',
            description: 'lorem ipsum dolor sit amet, consectetur adipiscing elit.',
          },
          {
            skills: ['contacts'],
            actions: ['empower'],
            modifier: 1,
            channel: 'tech',
            description: 'lorem ipsum dolor sit amet, consectetur adipiscing elit.',
          },
        ],
      },
    ],
    wounds: [
      {
        name: 'broken leg',
        description: 'leg broked',
        effects: [
          {
            skills: ['perform', 'insight'],
            actions: ['empower'],
            modifier: -5,
            channel: '',
          },
        ],
      },
    ],
    inventory: [
      {
        name: 'straw hat',
        description: 'lorem ipsum dolor sit amet, consectetur adipiscing elit.',
        effects: [
          {
            skills: ['persuade', 'operate'],
            actions: ['overcome'],
            modifier: 5,
            channel: '',
          },
        ],
      },
    ],
    modifiers: {},
  }),

  // Actions are for modifying state
  actions: {
    setCharacter(data) {
      this.character = data
    },
    addToInventory(item) {
      this.inventory.push(item)
    },
    updateSkill(skill, val) {
      this.skills[skill] = val
    },
    computeModifiers: (store) => {
      const { skills, aspects, wounds, inventory } = store
      const modifiers = {}

      // init modifiers
      const actions = ['attack', 'defend', 'overcome', 'empower']
      for (const skill in skills) {
        modifiers[skill] = { name: skill, sources: [] }
        actions.forEach((action) => {
          modifiers[skill][action] = skills[skill]
        })
      }

      const applyModifiers = (source, sourceName) => {
        source.forEach((item) => {
          item.effects.forEach((effect) => {
            effect.skills.forEach((skill) => {
              effect.actions.forEach((action) => {
                if (modifiers[skill] && modifiers[skill][action] !== undefined) {
                  modifiers[skill].sources.push(item)
                  modifiers[skill][action] += effect.modifier
                } else {
                  console.warn(`skipped ${skill} ${action} from ${sourceName} ${item.name}`)
                }
              })
            })
          })
        })
      }
      // Step 2: Apply effects from aspects, wounds, and inventory
      applyModifiers(aspects, 'aspect')
      applyModifiers(wounds, 'wound')
      applyModifiers(inventory, 'inventory')

      // Update the store's modifiers
      store.modifiers = { ...modifiers }
    },
  },

  // Getters are computed properties for state
  getters: {
    inventoryCount: (state) => state.inventory.length,
    characterName: (state) => state.name,
  },
})

<template>
  <div class="inventory">
    <h2>Inventory</h2>

    <table>
      <thead>
        <tr>
          <th>Quantity</th>
          <th>Name</th>
          <th>Type</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in items" :key="index">
          <td
            class="quantity-cell"
            @mouseenter="hoveredRow = index"
            @mouseleave="hoveredRow = null"
          >
            <span>{{ item.quantity }}</span>
            <div v-if="hoveredRow === index" class="quantity-adjust">
              <button @click="adjustQuantity(index, -1)">-</button>
              <button @click="adjustQuantity(index, 1)">+</button>
            </div>
          </td>
          <td>{{ item.name }}</td>
          <td>{{ item.type }}</td>
          <td>{{ item.description }}</td>
        </tr>
      </tbody>
    </table>

    <div class="add-item">
      <h3>Add Item</h3>
      <input v-model="newItem.name" placeholder="Name" />
      <input v-model="newItem.type" placeholder="Type" />
      <input v-model="newItem.description" placeholder="Description" />
      <input v-model.number="newItem.quantity" placeholder="Quantity" type="number" />
      <button @click="addItem">Add</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: [
        // Example items
        { quantity: 2, name: "Health Potion", type: "Consumable", description: "Restores health." },
        { quantity: 1, name: "Longsword", type: "Weapon", description: "A sharp blade." },
      ],
      newItem: { quantity: 0, name: "", type: "", description: "" },
      hoveredRow: null,
    };
  },
  methods: {
    addItem() {
      if (this.newItem.name && this.newItem.type && this.newItem.quantity >= 0) {
        this.items.push({ ...this.newItem });
        this.newItem = { quantity: 0, name: "", type: "", description: "" };
      } else {
        alert("Please fill out all fields and ensure quantity is non-negative.");
      }
    },
    adjustQuantity(index, amount) {
      const updatedQuantity = this.items[index].quantity + amount;
      if (updatedQuantity >= 0) {
        this.items[index].quantity = updatedQuantity;
      }
    },
  },
};
</script>

<style scoped>
.inventory {
  padding: 1rem;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 1rem;
}

th, td {
  padding: 0.5rem;
  border: 1px solid #ccc;
  text-align: left;
}

.quantity-cell {
  position: relative;
}

.quantity-adjust {
  position: absolute;
  top: 0;
  left: 0;
  display: flex;
  gap: 0.25rem;
  background: white;
  border: 1px solid #ccc;
  padding: 0.2rem;
}

.add-item {
  margin-top: 1rem;
}

.add-item input {
  margin: 0.5rem 0;
  padding: 0.5rem;
  width: calc(100% - 1rem);
}

.add-item button {
  padding: 0.5rem 1rem;
  margin-top: 0.5rem;
}
</style>

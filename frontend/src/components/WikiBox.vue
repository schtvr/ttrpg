<!-- WikiComponent.vue -->
<template>
  <div class="wiki">
    <h2>Wiki</h2>

    <div class="search-bar">
      <input v-model="searchQuery" @input="filterEntries" placeholder="Search the wiki..." />
    </div>

    <div class="wiki-list">
      <ul>
        <li v-for="entry in filteredEntries" :key="entry.id" @click="selectEntry(entry)">
          {{ entry.title }}
        </li>
      </ul>
    </div>

    <div class="wiki-editor" v-if="selectedEntry">
      <h3>Edit Entry: {{ selectedEntry.title }}</h3>
      <textarea v-model="selectedEntry.content" placeholder="Write in Markdown"></textarea>

      <label for="image-upload">Upload Image:</label>
      <input type="file" id="image-upload" @change="uploadImage" />

      <button @click="saveEntry">Save</button>
    </div>

    <div class="wiki-preview" v-if="selectedEntry">
      <h3>Preview</h3>
      <div v-html="renderMarkdown(selectedEntry.content)"></div>
    </div>
  </div>
</template>

<script>
import { marked } from 'marked';

export default {
  data() {
    return {
      entries: [
        // Example entries
        { id: 1, title: 'Story Overview', content: '## This is the story overview...' },
        { id: 2, title: 'Main Characters', content: '### Characters:\n- Hero\n- Villain' },
        { id: 3, title: 'Locations', content: '### Locations:\n- City of Dreams\n- Mystic Forest' },
      ],
      filteredEntries: [],
      selectedEntry: null,
      searchQuery: '',
    };
  },
  methods: {
    filterEntries() {
      this.filteredEntries = this.entries.filter(entry =>
        entry.title.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
    selectEntry(entry) {
      this.selectedEntry = entry;
    },
    saveEntry() {
      // Save logic (e.g., send to server)
      alert('Entry saved!');
    },
    uploadImage(event) {
      const file = event.target.files[0];
      if (file) {
        // Example: Upload to a server or handle locally
        const reader = new FileReader();
        reader.onload = () => {
          const imageUrl = reader.result;
          this.selectedEntry.content += `\n![Uploaded Image](${imageUrl})`;
        };
        reader.readAsDataURL(file);
      }
    },
    renderMarkdown(content) {
      return marked(content);
    },
  },
  mounted() {
    this.filteredEntries = this.entries;
  },
};
</script>

<style scoped>
.wiki {
  padding: 1rem;
}

.search-bar {
  margin-bottom: 1rem;
}

.wiki-list ul {
  list-style: none;
  padding: 0;
}

.wiki-list li {
  padding: 0.5rem;
  cursor: pointer;
  border-bottom: 1px solid #ccc;
}

.wiki-list li:hover {
  background-color: #f0f0f0;
}

.wiki-editor textarea {
  width: 100%;
  height: 200px;
  margin-bottom: 1rem;
}

.wiki-preview {
  margin-top: 1rem;
  padding: 1rem;
  border: 1px solid #ccc;
  background-color: #f9f9f9;
}
</style>

<!-- Explanation -->
<!--
This component is designed as a basic wiki system with the following features:

1. **Search**:
   - A search bar filters wiki entries based on the title.

2. **Markdown Editing**:
   - Entries can be written and edited using Markdown syntax.
   - The `marked` library is used to render Markdown into HTML.

3. **Image Upload**:
   - Users can upload images, which are converted to Base64 and appended to the entry.

4. **Preview**:
   - A live preview of the rendered Markdown is displayed beside the editor.

Replace the placeholder `entries` with actual data fetched from a database or API. Add server integration for saving and loading entries as needed.
-->

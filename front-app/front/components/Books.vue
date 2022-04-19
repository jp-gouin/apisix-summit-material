<template>
     <table
    class="table-auto text-sm border-0 overflow-auto h-8 sm:h-12 md:h-16 lg:h-20 xl:h-24 w-32 m-auto"
  >
    <thead>
      <tr class="border-b-2 border-gray-500">
        <th v-for="(column, index) in columns" :key="index" class="px-4 py-2">
          {{ column.label }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(book, index) in books" :key="index">
        <td class="border-0 px-4 py-2">{{ format(book).Title }}</td>
        <td class="border-0 px-4 py-2">{{ format(book).AuthorName }}</td>
      </tr>
    </tbody>
  </table>
  
</template>
<script>
export default {
  props: {
    books: Array,
  },
  methods: {
    format(book) {
      return JSON.parse(book);
      //return book;
    },
    getAuthorName(book){
      let myvar = this.format(book).Author
      return myvar
    }
  },
  data() {
    return {
      columns: [
        {
          label: "Title",
          jsonId: "",
        },
        {
          label: "Author",
          jsonId: "",
        }
      ],
    };
  },
};
</script>
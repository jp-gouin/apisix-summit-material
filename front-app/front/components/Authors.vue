
<template>
  <div>
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
        <tr v-for="(author, index) in authors" :key="index+'author'">
          <td class="border-0 px-4 py-2">{{ format(author).Name }}</td>
          <td class="border-0 px-4 py-2">{{ format(author).LastName }}</td>
          <td class="border-0 px-4 py-2">{{ format(author).Number }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script>
export default {
  props: {
    authors: Array,
  },
  methods: {
    format(author) {
      console.log(JSON.parse(author))
      return JSON.parse(author);
      //return author;
    },
  },
  data() {
    return {
      columns: [
        {
          label: "Name",
          jsonId: "",
        },
        {
          label: "Lastname",
          jsonId: "",
        },{
          label: "Number",
          jsonId: "",
        },
      ],
    };
  },
};
</script>
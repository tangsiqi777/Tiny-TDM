<script setup>
const props = defineProps(['total'])
</script>

<template>

  <div class="table-page-c">
    <a-pagination :total="props.total" simple/>
  </div>

</template>

<style scoped>

</style>
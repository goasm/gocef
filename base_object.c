#include "base_object.h"

#include <stdlib.h>
#include "include/capi/cef_base_capi.h"

typedef struct _ref_counted_impl_t {
  size_t count;
} ref_counted_impl_t;

static inline ref_counted_impl_t* ref_counted(cef_base_ref_counted_t* base) {
  uint8_t* p = (uint8_t*)base;
  return (ref_counted_impl_t*)(p + base->size);
}

static void add_ref_impl(cef_base_ref_counted_t* self) {
  ref_counted_impl_t* refcnt = ref_counted(self);
  refcnt->count = refcnt->count + 1;
}

static int release_impl(cef_base_ref_counted_t* self) {
  ref_counted_impl_t* refcnt = ref_counted(self);
  refcnt->count = refcnt->count - 1;
  if (refcnt->count == 0) {
    free(self);
  }
}

static int has_one_ref_impl(cef_base_ref_counted_t* self) {
  ref_counted_impl_t* refcnt = ref_counted(self);
  return refcnt->count == 1;
}

static int has_at_least_one_ref_impl(cef_base_ref_counted_t* self) {
  ref_counted_impl_t* refcnt = ref_counted(self);
  return refcnt->count >= 1;
}

void* gocef_new_impl(size_t size) {
  void* obj = calloc(1, size);
  cef_base_ref_counted_t* base = (cef_base_ref_counted_t*)obj;
  base->size = size;
  base->add_ref = add_ref_impl;
  base->release = release_impl;
  base->has_one_ref = has_one_ref_impl;
  base->has_at_least_one_ref = has_at_least_one_ref_impl;
  ref_counted(base)->count = 1;
  return obj;
}

void gocef_add_ref_impl(void* ptr) {
  cef_base_ref_counted_t* base = (cef_base_ref_counted_t*)ptr;
  base->add_ref(base);
}

int gocef_release_impl(void* ptr) {
  cef_base_ref_counted_t* base = (cef_base_ref_counted_t*)ptr;
  return base->release(base);
}

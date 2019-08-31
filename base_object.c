#include "base_object.h"

#include <stdlib.h>
#include "include/capi/cef_base_capi.h"

static void add_ref_impl(cef_base_ref_counted_t* self) {
}

static int release_impl(cef_base_ref_counted_t* self) {
}

static int has_one_ref_impl(cef_base_ref_counted_t* self) {
}

void* gocef_new(size_t size) {
  void* obj = calloc(1, size);
  cef_base_ref_counted_t* base = (cef_base_ref_counted_t*)obj;
  base->size = size;
  base->add_ref = add_ref_impl;
  base->release = release_impl;
  base->has_one_ref = has_one_ref_impl;
  return obj;
}

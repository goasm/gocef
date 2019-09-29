#include "type_conv.h"

#include <stdlib.h>

void gocef_set_utf8_impl(cef_string_utf8_t* dst, char* str, size_t len) {
  cef_string_utf8_set(str, len, dst, 1);
  free(str);
}

void gocef_set_utf16_impl(cef_string_utf16_t* dst, char* str, size_t len) {
  cef_string_utf8_to_utf16(str, len, dst);
  free(str);
}

#include "type_conv.h"

#include <stdlib.h>

cef_string_utf8_t* gocef_to_utf8_impl(char* str, size_t len) {
  cef_string_utf8_t* p = calloc(1, sizeof(cef_string_utf8_t));
  p->str = str;
  p->length = len;
  return p;
}

cef_string_utf16_t* gocef_to_utf16_impl(char* str, size_t len) {
  cef_string_utf16_t* p = calloc(1, sizeof(cef_string_utf16_t));
  cef_string_utf8_to_utf16(str, len, p);
  return p;
}

#include "type_conv.h"

cef_string_utf8_t gocef_to_utf8_impl(char* str, size_t len) {
  cef_string_utf8_t s;
  s.str = str;
  s.length = len;
  return s;
}

cef_string_utf16_t gocef_to_utf16_impl(char* str, size_t len) {
  cef_string_utf16_t s;
  cef_string_utf8_to_utf16(str, len, &s);
  return s;
}

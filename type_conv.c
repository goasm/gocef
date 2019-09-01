#include "type_conv.h"

int gocef_set_string_utf8(_GoString_ src, cef_string_utf8_t* out) {
  const char* str = _GoStringPtr(src);
  size_t length = _GoStringLen(src);
  out->str = (char*)str;
  out->length = length;
  return length;
}

int gocef_set_string_utf16(_GoString_ src, cef_string_utf16_t* out) {
  const char* str = _GoStringPtr(src);
  size_t length = _GoStringLen(src);
  return cef_string_utf8_to_utf16(str, length, out);
}

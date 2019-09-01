// gocef type conversion utils
#ifndef GOCEF_TYPE_CONV_H_
#define GOCEF_TYPE_CONV_H_
#pragma once

#include "include/internal/cef_string.h"

int gocef_set_string_utf8(_GoString_ src, cef_string_utf8_t* out);
int gocef_set_string_utf16(_GoString_ src, cef_string_utf16_t* out);

#endif  // GOCEF_TYPE_CONV_H_

// gocef type conversion
#ifndef GOCEF_TYPE_CONV_H_
#define GOCEF_TYPE_CONV_H_
#pragma once

#include "include/internal/cef_string.h"

cef_string_utf8_t gocef_to_utf8_impl(char* str, size_t len);
cef_string_utf16_t gocef_to_utf16_impl(char* str, size_t len);

#endif  // GOCEF_TYPE_CONV_H_

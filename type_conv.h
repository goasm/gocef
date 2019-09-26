// gocef type conversion
#ifndef GOCEF_TYPE_CONV_H_
#define GOCEF_TYPE_CONV_H_
#pragma once

#include "include/internal/cef_string.h"

void gocef_set_utf8_impl(cef_string_utf8_t* dst, char* str, size_t len);
void gocef_set_utf16_impl(cef_string_utf16_t* dst, char* str, size_t len);

#endif  // GOCEF_TYPE_CONV_H_

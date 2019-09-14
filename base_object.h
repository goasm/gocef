// gocef base object
#ifndef GOCEF_BASE_OBJECT_H_
#define GOCEF_BASE_OBJECT_H_
#pragma once

#include <stddef.h>

void* gocef_new_impl(size_t size);
void gocef_add_ref_impl(void* ptr);
int gocef_release_impl(void* ptr);

#endif  // GOCEF_BASE_OBJECT_H_

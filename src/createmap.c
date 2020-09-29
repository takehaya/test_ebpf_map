#define KBUILD_MODNAME "createmap"

#include <linux/bpf.h>
#include "bpf_helpers.h"

struct inner_map
{
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __uint(max_entries, 1);
    __type(key, int);
    __type(value, int);
} inner_map SEC(".maps");

struct outer_arr
{
    __uint(type, BPF_MAP_TYPE_ARRAY_OF_MAPS);
    __uint(max_entries, 1);
    __uint(key_size, sizeof(int));
    __uint(value_size, sizeof(int));
    __array(values, struct inner_map);
} outer_arr SEC(".maps") = {
    .values = {&inner_map},
};

struct outer_hash
{
    __uint(type, BPF_MAP_TYPE_HASH_OF_MAPS);
    __uint(max_entries, 5);
    __uint(key_size, sizeof(int));
    __array(values, struct inner_map);
} outer_hash SEC(".maps") = {
    .values = {&inner_map},
};

// Copyright 2023 Ingonyama
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	
// Code generated by Ingonyama DO NOT EDIT

#include <cuda.h>
#include <stdbool.h>
// projective.h

#ifdef __cplusplus
extern "C" {
#endif

typedef struct BW6761_projective_t BW6761_projective_t;
typedef struct BW6761_g2_projective_t BW6761_g2_projective_t;
typedef struct BW6761_affine_t BW6761_affine_t;
typedef struct BW6761_g2_affine_t BW6761_g2_affine_t;
typedef struct BW6761_scalar_t BW6761_scalar_t;

bool projective_is_on_curve_bw6761(BW6761_projective_t* point1);

int random_scalar_bw6761(BW6761_scalar_t* out); 
int random_projective_bw6761(BW6761_projective_t* out);
BW6761_projective_t* projective_zero_bw6761();
int projective_to_affine_bw6761(BW6761_affine_t* out, BW6761_projective_t* point1);
int projective_from_affine_bw6761(BW6761_projective_t* out, BW6761_affine_t* point1);

int random_g2_projective_bw6761(BW6761_g2_projective_t* out);
int g2_projective_to_affine_bw6761(BW6761_g2_affine_t* out, BW6761_g2_projective_t* point1);
int g2_projective_from_affine_bw6761(BW6761_g2_projective_t* out, BW6761_g2_affine_t* point1);
bool g2_projective_is_on_curve_bw6761(BW6761_g2_projective_t* point1);

bool eq_bw6761(BW6761_projective_t* point1, BW6761_projective_t* point2);
bool eq_g2_bw6761(BW6761_g2_projective_t* point1, BW6761_g2_projective_t* point2);

#ifdef __cplusplus
}
#endif

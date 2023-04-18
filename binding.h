#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

typedef void* GPTJHandle;
typedef void* ContextHandle;

typedef int (*Callback)(const char*);

struct GPTJContext {
        GPTJHandle handle;
        ContextHandle ctx;
};

struct GPTJContext* go_gptj_load(char* modelFullPath);

void go_gptj_prompt(struct GPTJContext ctx, char* prompt, Callback responseCallback, 
                int32_t n_predict, int32_t top_k, float top_p, float temp, int32_t n_batch); 

#ifdef __cplusplus
}
#endif

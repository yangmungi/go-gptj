#include <string>
#include <fstream>

#include "binding.h"
#include "gpt4all-chat/gptj.h"

GPTJContext go_gptj_load(char* modelFullPath) {
       std::string fullPath (modelFullPath);
       auto fin = std::ifstream(fullPath, std::ios::binary);

       GPTJ* m =  new GPTJ();
       m->loadModel(fullPath, fin);

       LLModel::PromptContext llmctx ();

       void* ctx = reinterpret_cast<void*>(&llmctx);

       GPTJContext gptjCtx = {m, ctx};

       return gptjCtx;
}

void go_gptj_prompt(GPTJContext ctx, char* prompt, Callback responseCallback,
                int32_t n_predict, int32_t top_k, float top_p, float temp, int32_t n_batch) {

        LLModel::PromptContext llmCtx = reinterpret_cast<LLModel::PromptContext>(ctx.ctx);
        std::string stdPrompt (prompt);

        ctx.handle->prompt(stdPrompt, responseCallback, llmCtx, n_predict, top_k, top_p, temp, n_batch);
}

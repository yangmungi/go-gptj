#include "binding.h"
#include "gpt4all-chat/gptj.h"

GPTJHandle load() {
        return new GPTJ();
}

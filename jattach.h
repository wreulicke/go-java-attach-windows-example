
typedef void (*CallbackFn)(uintptr_t ch, char* buf);
typedef void (*CloseCallbackFn)(uintptr_t ch, int r);

typedef struct {
  CallbackFn callback;
  CloseCallbackFn closeCallback;
} Callbacks;

extern int inject_thread(int pid, char* pipeName, int argc, char** argv);
extern int attach(uintptr_t ch, Callbacks cbs, int pid, char* pipeName, int argc, char** argv);
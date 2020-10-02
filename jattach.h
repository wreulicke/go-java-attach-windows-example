
typedef void (*CallbackFn)(void* ch, char* buf);
typedef void (*CloseCallbackFn)(void* ch, int r);

typedef struct {
  CallbackFn callback;
  CloseCallbackFn closeCallback;
} Callbacks;

extern int inject_thread(int pid, char* pipeName, int argc, char** argv);
extern int attach(void* ch, Callbacks cbs, int pid, char* pipeName, int argc, char** argv);
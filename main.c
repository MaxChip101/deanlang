/* single_byte_rw.c
 *
 * No #includes. Simple syscall wrappers for single-byte read/write.
 * Supports: x86_64, i386, aarch64, arm (32-bit).
 *
 * Build (on the target machine/arch):
 *   gcc -O2 -nostdlib -static -s -o single_byte_rw single_byte_rw.c
 * or just:
 *   gcc -O2 -o single_byte_rw single_byte_rw.c
 *
 * (You don't need -nostdlib; it's optional. This file doesn't include headers.)
 *
 * Behavior:
 *  - read1(fd, &c) reads up to 1 byte from fd into c, returns number of bytes read (0 for EOF, negative on error)
 *  - write1(fd, &c) writes 1 byte, returns number of bytes written (or negative on error)
 */

typedef long long ll;
typedef long lr;       /* generic long for return values */

/* Prototype */
lr read1(int fd, void *buf);
lr write1(int fd, const void *buf);

/* Architecture-specific implementations */
#if defined(__x86_64__)

/* Linux x86_64: syscall in rax, arg1=rdi, arg2=rsi, arg3=rdx */
lr read1(int fd, void *buf) {
    lr ret;
    register long rax __asm__("rax") = 0;   /* __NR_read = 0 */
    register long rdi __asm__("rdi") = fd;
    register void *rsi __asm__("rsi") = buf;
    register long rdx __asm__("rdx") = 1;
    __asm__ volatile (
        "syscall"
        : "=r"(rax)
        : "r"(rax), "r"(rdi), "r"(rsi), "r"(rdx)
        : "rcx", "r11", "memory"
    );
    ret = (lr)rax;
    return ret;
}

lr write1(int fd, const void *buf) {
    lr ret;
    register long rax __asm__("rax") = 1;   /* __NR_write = 1 */
    register long rdi __asm__("rdi") = fd;
    register const void *rsi __asm__("rsi") = buf;
    register long rdx __asm__("rdx") = 1;
    __asm__ volatile (
        "syscall"
        : "=r"(rax)
        : "r"(rax), "r"(rdi), "r"(rsi), "r"(rdx)
        : "rcx", "r11", "memory"
    );
    ret = (lr)rax;
    return ret;
}

#elif defined(__i386__)

/* Linux x86 (32-bit) using int 0x80: eax = syscall, ebx, ecx, edx args */
lr read1(int fd, void *buf) {
    lr ret;
    int syscall_no = 3; /* __NR_read = 3 on i386 */
    __asm__ volatile (
        "int $0x80"
        : "=a"(ret)
        : "0"(syscall_no), "b"(fd), "c"(buf), "d"(1)
        : "memory"
    );
    return ret;
}

lr write1(int fd, const void *buf) {
    lr ret;
    int syscall_no = 4; /* __NR_write = 4 on i386 */
    __asm__ volatile (
        "int $0x80"
        : "=a"(ret)
        : "0"(syscall_no), "b"(fd), "c"(buf), "d"(1)
        : "memory"
    );
    return ret;
}

#elif defined(__aarch64__)

/* aarch64: syscall number in x8, args x0..x2, svc #0.
 * __NR_read = 63, __NR_write = 64 on aarch64 linux
 */
lr read1(int fd, void *buf) {
    register long x8 __asm__("x8") = 63; /* __NR_read */
    register long x0 __asm__("x0") = fd;
    register void *x1 __asm__("x1") = buf;
    register long x2 __asm__("x2") = 1;
    __asm__ volatile (
        "svc #0"
        : "+r"(x0)
        : "r"(x8), "r"(x1), "r"(x2)
        : "memory"
    );
    return (lr)x0;
}

lr write1(int fd, const void *buf) {
    register long x8 __asm__("x8") = 64; /* __NR_write */
    register long x0 __asm__("x0") = fd;
    register const void *x1 __asm__("x1") = buf;
    register long x2 __asm__("x2") = 1;
    __asm__ volatile (
        "svc #0"
        : "+r"(x0)
        : "r"(x8), "r"(x1), "r"(x2)
        : "memory"
    );
    return (lr)x0;
}

#elif defined(__arm__)

/* 32-bit ARM EABI: syscall via swi 0 (or svc 0), r7=syscall number.
 * __NR_read = 3, __NR_write = 4
 */
lr read1(int fd, void *buf) {
    lr ret;
    register int r7 __asm__("r7") = 3; /* __NR_read */
    register int r0 __asm__("r0") = fd;
    register void *r1 __asm__("r1") = buf;
    register int r2 __asm__("r2") = 1;
    __asm__ volatile (
        "svc 0"
        : "+r"(r0)
        : "r"(r7), "r"(r1), "r"(r2)
        : "memory"
    );
    ret = (lr)r0;
    return ret;
}

lr write1(int fd, const void *buf) {
    lr ret;
    register int r7 __asm__("r7") = 4; /* __NR_write */
    register int r0 __asm__("r0") = fd;
    register const void *r1 __asm__("r1") = buf;
    register int r2 __asm__("r2") = 1;
    __asm__ volatile (
        "svc 0"
        : "+r"(r0)
        : "r"(r7), "r"(r1), "r"(r2)
        : "memory"
    );
    ret = (lr)r0;
    return ret;
}

#else
#error "Unsupported architecture. Supported: x86_64, i386, aarch64, arm."
#endif

/* Example program: read one byte from stdin (fd 0) and write it to stdout (fd 1).
 * Returns syscall exit status as the process exit code (we don't call exit()).
 *
 * NOTE: We can't call exit() portably without includes / libc; returning value from
 * main becomes the process exit code when linked normally.
 */

void _start (void) {
    unsigned char c = 0;
    lr r;

    r = read1(0, &c); /* fd 0 = stdin */
    if (r <= 0) {
        /* EOF or error: return 1 for error/0 for EOF (we return 1 on error) */
        if (r < 0) exit(1);
        exit(0);
    }

    r = write1(1, &c); /* fd 1 = stdout */
    if (r <= 0) exit(1);

    exit(0);
}

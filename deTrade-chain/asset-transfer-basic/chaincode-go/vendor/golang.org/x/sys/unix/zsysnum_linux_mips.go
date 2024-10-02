// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/mips/include /tmp/mips/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build mips && linux

package unix

const (
	SYS_SYSCALL                      = 4000
	SYS_EXIT                         = 4001
	SYS_FORK                         = 4002
	SYS_READ                         = 4003
	SYS_WRITE                        = 4004
	SYS_OPEN                         = 4005
	SYS_CLOSE                        = 4006
	SYS_WAITPID                      = 4007
	SYS_CREAT                        = 4008
	SYS_LINK                         = 4009
	SYS_UNLINK                       = 4010
	SYS_EXECVE                       = 4011
	SYS_CHDIR                        = 4012
	SYS_TIME                         = 4013
	SYS_MKNOD                        = 4014
	SYS_CHMOD                        = 4015
	SYS_LCHOWN                       = 4016
	SYS_BREAK                        = 4017
	SYS_UNUSED18                     = 4018
	SYS_LSEEK                        = 4019
	SYS_GETPID                       = 4020
	SYS_MOUNT                        = 4021
	SYS_UMOUNT                       = 4022
	SYS_SETUID                       = 4023
	SYS_GETUID                       = 4024
	SYS_STIME                        = 4025
	SYS_PTRACE                       = 4026
	SYS_ALARM                        = 4027
	SYS_UNUSED28                     = 4028
	SYS_PAUSE                        = 4029
	SYS_UTIME                        = 4030
	SYS_STTY                         = 4031
	SYS_GTTY                         = 4032
	SYS_ACCESS                       = 4033
	SYS_NICE                         = 4034
	SYS_FTIME                        = 4035
	SYS_SYNC                         = 4036
	SYS_KILL                         = 4037
	SYS_RENAME                       = 4038
	SYS_MKDIR                        = 4039
	SYS_RMDIR                        = 4040
	SYS_DUP                          = 4041
	SYS_PIPE                         = 4042
	SYS_TIMES                        = 4043
	SYS_PROF                         = 4044
	SYS_BRK                          = 4045
	SYS_SETGID                       = 4046
	SYS_GETGID                       = 4047
	SYS_SIGNAL                       = 4048
	SYS_GETEUID                      = 4049
	SYS_GETEGID                      = 4050
	SYS_ACCT                         = 4051
	SYS_UMOUNT2                      = 4052
	SYS_LOCK                         = 4053
	SYS_IOCTL                        = 4054
	SYS_FCNTL                        = 4055
	SYS_MPX                          = 4056
	SYS_SETPGID                      = 4057
	SYS_ULIMIT                       = 4058
	SYS_UNUSED59                     = 4059
	SYS_UMASK                        = 4060
	SYS_CHROOT                       = 4061
	SYS_USTAT                        = 4062
	SYS_DUP2                         = 4063
	SYS_GETPPID                      = 4064
	SYS_GETPGRP                      = 4065
	SYS_SETSID                       = 4066
	SYS_SIGACTION                    = 4067
	SYS_SGETMASK                     = 4068
	SYS_SSETMASK                     = 4069
	SYS_SETREUID                     = 4070
	SYS_SETREGID                     = 4071
	SYS_SIGSUSPEND                   = 4072
	SYS_SIGPENDING                   = 4073
	SYS_SETHOSTNAME                  = 4074
	SYS_SETRLIMIT                    = 4075
	SYS_GETRLIMIT                    = 4076
	SYS_GETRUSAGE                    = 4077
	SYS_GETTIMEOFDAY                 = 4078
	SYS_SETTIMEOFDAY                 = 4079
	SYS_GETGROUPS                    = 4080
	SYS_SETGROUPS                    = 4081
	SYS_RESERVED82                   = 4082
	SYS_SYMLINK                      = 4083
	SYS_UNUSED84                     = 4084
	SYS_READLINK                     = 4085
	SYS_USELIB                       = 4086
	SYS_SWAPON                       = 4087
	SYS_REBOOT                       = 4088
	SYS_READDIR                      = 4089
	SYS_MMAP                         = 4090
	SYS_MUNMAP                       = 4091
	SYS_TRUNCATE                     = 4092
	SYS_FTRUNCATE                    = 4093
	SYS_FCHMOD                       = 4094
	SYS_FCHOWN                       = 4095
	SYS_GETPRIORITY                  = 4096
	SYS_SETPRIORITY                  = 4097
	SYS_PROFIL                       = 4098
	SYS_STATFS                       = 4099
	SYS_FSTATFS                      = 4100
	SYS_IOPERM                       = 4101
	SYS_SOCKETCALL                   = 4102
	SYS_SYSLOG                       = 4103
	SYS_SETITIMER                    = 4104
	SYS_GETITIMER                    = 4105
	SYS_STAT                         = 4106
	SYS_LSTAT                        = 4107
	SYS_FSTAT                        = 4108
	SYS_UNUSED109                    = 4109
	SYS_IOPL                         = 4110
	SYS_VHANGUP                      = 4111
	SYS_IDLE                         = 4112
	SYS_VM86                         = 4113
	SYS_WAIT4                        = 4114
	SYS_SWAPOFF                      = 4115
	SYS_SYSINFO                      = 4116
	SYS_IPC                          = 4117
	SYS_FSYNC                        = 4118
	SYS_SIGRETURN                    = 4119
	SYS_CLONE                        = 4120
	SYS_SETDOMAINNAME                = 4121
	SYS_UNAME                        = 4122
	SYS_MODIFY_LDT                   = 4123
	SYS_ADJTIMEX                     = 4124
	SYS_MPROTECT                     = 4125
	SYS_SIGPROCMASK                  = 4126
	SYS_CREATE_MODULE                = 4127
	SYS_INIT_MODULE                  = 4128
	SYS_DELETE_MODULE                = 4129
	SYS_GET_KERNEL_SYMS              = 4130
	SYS_QUOTACTL                     = 4131
	SYS_GETPGID                      = 4132
	SYS_FCHDIR                       = 4133
	SYS_BDFLUSH                      = 4134
	SYS_SYSFS                        = 4135
	SYS_PERSONALITY                  = 4136
	SYS_AFS_SYSCALL                  = 4137
	SYS_SETFSUID                     = 4138
	SYS_SETFSGID                     = 4139
	SYS__LLSEEK                      = 4140
	SYS_GETDENTS                     = 4141
	SYS__NEWSELECT                   = 4142
	SYS_FLOCK                        = 4143
	SYS_MSYNC                        = 4144
	SYS_READV                        = 4145
	SYS_WRITEV                       = 4146
	SYS_CACHEFLUSH                   = 4147
	SYS_CACHECTL                     = 4148
	SYS_SYSMIPS                      = 4149
	SYS_UNUSED150                    = 4150
	SYS_GETSID                       = 4151
	SYS_FDATASYNC                    = 4152
	SYS__SYSCTL                      = 4153
	SYS_MLOCK                        = 4154
	SYS_MUNLOCK                      = 4155
	SYS_MLOCKALL                     = 4156
	SYS_MUNLOCKALL                   = 4157
	SYS_SCHED_SETPARAM               = 4158
	SYS_SCHED_GETPARAM               = 4159
	SYS_SCHED_SETSCHEDULER           = 4160
	SYS_SCHED_GETSCHEDULER           = 4161
	SYS_SCHED_YIELD                  = 4162
	SYS_SCHED_GET_PRIORITY_MAX       = 4163
	SYS_SCHED_GET_PRIORITY_MIN       = 4164
	SYS_SCHED_RR_GET_INTERVAL        = 4165
	SYS_NANOSLEEP                    = 4166
	SYS_MREMAP                       = 4167
	SYS_ACCEPT                       = 4168
	SYS_BIND                         = 4169
	SYS_CONNECT                      = 4170
	SYS_GETPEERNAME                  = 4171
	SYS_GETSOCKNAME                  = 4172
	SYS_GETSOCKOPT                   = 4173
	SYS_LISTEN                       = 4174
	SYS_RECV                         = 4175
	SYS_RECVFROM                     = 4176
	SYS_RECVMSG                      = 4177
	SYS_SEND                         = 4178
	SYS_SENDMSG                      = 4179
	SYS_SENDTO                       = 4180
	SYS_SETSOCKOPT                   = 4181
	SYS_SHUTDOWN                     = 4182
	SYS_SOCKET                       = 4183
	SYS_SOCKETPAIR                   = 4184
	SYS_SETRESUID                    = 4185
	SYS_GETRESUID                    = 4186
	SYS_QUERY_MODULE                 = 4187
	SYS_POLL                         = 4188
	SYS_NFSSERVCTL                   = 4189
	SYS_SETRESGID                    = 4190
	SYS_GETRESGID                    = 4191
	SYS_PRCTL                        = 4192
	SYS_RT_SIGRETURN                 = 4193
	SYS_RT_SIGACTION                 = 4194
	SYS_RT_SIGPROCMASK               = 4195
	SYS_RT_SIGPENDING                = 4196
	SYS_RT_SIGTIMEDWAIT              = 4197
	SYS_RT_SIGQUEUEINFO              = 4198
	SYS_RT_SIGSUSPEND                = 4199
	SYS_PREAD64                      = 4200
	SYS_PWRITE64                     = 4201
	SYS_CHOWN                        = 4202
	SYS_GETCWD                       = 4203
	SYS_CAPGET                       = 4204
	SYS_CAPSET                       = 4205
	SYS_SIGALTSTACK                  = 4206
	SYS_SENDFILE                     = 4207
	SYS_GETPMSG                      = 4208
	SYS_PUTPMSG                      = 4209
	SYS_MMAP2                        = 4210
	SYS_TRUNCATE64                   = 4211
	SYS_FTRUNCATE64                  = 4212
	SYS_STAT64                       = 4213
	SYS_LSTAT64                      = 4214
	SYS_FSTAT64                      = 4215
	SYS_PIVOT_ROOT                   = 4216
	SYS_MINCORE                      = 4217
	SYS_MADVISE                      = 4218
	SYS_GETDENTS64                   = 4219
	SYS_FCNTL64                      = 4220
	SYS_RESERVED221                  = 4221
	SYS_GETTID                       = 4222
	SYS_READAHEAD                    = 4223
	SYS_SETXATTR                     = 4224
	SYS_LSETXATTR                    = 4225
	SYS_FSETXATTR                    = 4226
	SYS_GETXATTR                     = 4227
	SYS_LGETXATTR                    = 4228
	SYS_FGETXATTR                    = 4229
	SYS_LISTXATTR                    = 4230
	SYS_LLISTXATTR                   = 4231
	SYS_FLISTXATTR                   = 4232
	SYS_REMOVEXATTR                  = 4233
	SYS_LREMOVEXATTR                 = 4234
	SYS_FREMOVEXATTR                 = 4235
	SYS_TKILL                        = 4236
	SYS_SENDFILE64                   = 4237
	SYS_FUTEX                        = 4238
	SYS_SCHED_SETAFFINITY            = 4239
	SYS_SCHED_GETAFFINITY            = 4240
	SYS_IO_SETUP                     = 4241
	SYS_IO_DESTROY                   = 4242
	SYS_IO_GETEVENTS                 = 4243
	SYS_IO_SUBMIT                    = 4244
	SYS_IO_CANCEL                    = 4245
	SYS_EXIT_GROUP                   = 4246
	SYS_LOOKUP_DCOOKIE               = 4247
	SYS_EPOLL_CREATE                 = 4248
	SYS_EPOLL_CTL                    = 4249
	SYS_EPOLL_WAIT                   = 4250
	SYS_REMAP_FILE_PAGES             = 4251
	SYS_SET_TID_ADDRESS              = 4252
	SYS_RESTART_SYSCALL              = 4253
	SYS_FADVISE64                    = 4254
	SYS_STATFS64                     = 4255
	SYS_FSTATFS64                    = 4256
	SYS_TIMER_CREATE                 = 4257
	SYS_TIMER_SETTIME                = 4258
	SYS_TIMER_GETTIME                = 4259
	SYS_TIMER_GETOVERRUN             = 4260
	SYS_TIMER_DELETE                 = 4261
	SYS_CLOCK_SETTIME                = 4262
	SYS_CLOCK_GETTIME                = 4263
	SYS_CLOCK_GETRES                 = 4264
	SYS_CLOCK_NANOSLEEP              = 4265
	SYS_TGKILL                       = 4266
	SYS_UTIMES                       = 4267
	SYS_MBIND                        = 4268
	SYS_GET_MEMPOLICY                = 4269
	SYS_SET_MEMPOLICY                = 4270
	SYS_MQ_OPEN                      = 4271
	SYS_MQ_UNLINK                    = 4272
	SYS_MQ_TIMEDSEND                 = 4273
	SYS_MQ_TIMEDRECEIVE              = 4274
	SYS_MQ_NOTIFY                    = 4275
	SYS_MQ_GETSETATTR                = 4276
	SYS_VSERVER                      = 4277
	SYS_WAITID                       = 4278
	SYS_ADD_KEY                      = 4280
	SYS_REQUEST_KEY                  = 4281
	SYS_KEYCTL                       = 4282
	SYS_SET_THREAD_AREA              = 4283
	SYS_INOTIFY_INIT                 = 4284
	SYS_INOTIFY_ADD_WATCH            = 4285
	SYS_INOTIFY_RM_WATCH             = 4286
	SYS_MIGRATE_PAGES                = 4287
	SYS_OPENAT                       = 4288
	SYS_MKDIRAT                      = 4289
	SYS_MKNODAT                      = 4290
	SYS_FCHOWNAT                     = 4291
	SYS_FUTIMESAT                    = 4292
	SYS_FSTATAT64                    = 4293
	SYS_UNLINKAT                     = 4294
	SYS_RENAMEAT                     = 4295
	SYS_LINKAT                       = 4296
	SYS_SYMLINKAT                    = 4297
	SYS_READLINKAT                   = 4298
	SYS_FCHMODAT                     = 4299
	SYS_FACCESSAT                    = 4300
	SYS_PSELECT6                     = 4301
	SYS_PPOLL                        = 4302
	SYS_UNSHARE                      = 4303
	SYS_SPLICE                       = 4304
	SYS_SYNC_FILE_RANGE              = 4305
	SYS_TEE                          = 4306
	SYS_VMSPLICE                     = 4307
	SYS_MOVE_PAGES                   = 4308
	SYS_SET_ROBUST_LIST              = 4309
	SYS_GET_ROBUST_LIST              = 4310
	SYS_KEXEC_LOAD                   = 4311
	SYS_GETCPU                       = 4312
	SYS_EPOLL_PWAIT                  = 4313
	SYS_IOPRIO_SET                   = 4314
	SYS_IOPRIO_GET                   = 4315
	SYS_UTIMENSAT                    = 4316
	SYS_SIGNALFD                     = 4317
	SYS_TIMERFD                      = 4318
	SYS_EVENTFD                      = 4319
	SYS_FALLOCATE                    = 4320
	SYS_TIMERFD_CREATE               = 4321
	SYS_TIMERFD_GETTIME              = 4322
	SYS_TIMERFD_SETTIME              = 4323
	SYS_SIGNALFD4                    = 4324
	SYS_EVENTFD2                     = 4325
	SYS_EPOLL_CREATE1                = 4326
	SYS_DUP3                         = 4327
	SYS_PIPE2                        = 4328
	SYS_INOTIFY_INIT1                = 4329
	SYS_PREADV                       = 4330
	SYS_PWRITEV                      = 4331
	SYS_RT_TGSIGQUEUEINFO            = 4332
	SYS_PERF_EVENT_OPEN              = 4333
	SYS_ACCEPT4                      = 4334
	SYS_RECVMMSG                     = 4335
	SYS_FANOTIFY_INIT                = 4336
	SYS_FANOTIFY_MARK                = 4337
	SYS_PRLIMIT64                    = 4338
	SYS_NAME_TO_HANDLE_AT            = 4339
	SYS_OPEN_BY_HANDLE_AT            = 4340
	SYS_CLOCK_ADJTIME                = 4341
	SYS_SYNCFS                       = 4342
	SYS_SENDMMSG                     = 4343
	SYS_SETNS                        = 4344
	SYS_PROCESS_VM_READV             = 4345
	SYS_PROCESS_VM_WRITEV            = 4346
	SYS_KCMP                         = 4347
	SYS_FINIT_MODULE                 = 4348
	SYS_SCHED_SETATTR                = 4349
	SYS_SCHED_GETATTR                = 4350
	SYS_RENAMEAT2                    = 4351
	SYS_SECCOMP                      = 4352
	SYS_GETRANDOM                    = 4353
	SYS_MEMFD_CREATE                 = 4354
	SYS_BPF                          = 4355
	SYS_EXECVEAT                     = 4356
	SYS_USERFAULTFD                  = 4357
	SYS_MEMBARRIER                   = 4358
	SYS_MLOCK2                       = 4359
	SYS_COPY_FILE_RANGE              = 4360
	SYS_PREADV2                      = 4361
	SYS_PWRITEV2                     = 4362
	SYS_PKEY_MPROTECT                = 4363
	SYS_PKEY_ALLOC                   = 4364
	SYS_PKEY_FREE                    = 4365
	SYS_STATX                        = 4366
	SYS_RSEQ                         = 4367
	SYS_IO_PGETEVENTS                = 4368
	SYS_SEMGET                       = 4393
	SYS_SEMCTL                       = 4394
	SYS_SHMGET                       = 4395
	SYS_SHMCTL                       = 4396
	SYS_SHMAT                        = 4397
	SYS_SHMDT                        = 4398
	SYS_MSGGET                       = 4399
	SYS_MSGSND                       = 4400
	SYS_MSGRCV                       = 4401
	SYS_MSGCTL                       = 4402
	SYS_CLOCK_GETTIME64              = 4403
	SYS_CLOCK_SETTIME64              = 4404
	SYS_CLOCK_ADJTIME64              = 4405
	SYS_CLOCK_GETRES_TIME64          = 4406
	SYS_CLOCK_NANOSLEEP_TIME64       = 4407
	SYS_TIMER_GETTIME64              = 4408
	SYS_TIMER_SETTIME64              = 4409
	SYS_TIMERFD_GETTIME64            = 4410
	SYS_TIMERFD_SETTIME64            = 4411
	SYS_UTIMENSAT_TIME64             = 4412
	SYS_PSELECT6_TIME64              = 4413
	SYS_PPOLL_TIME64                 = 4414
	SYS_IO_PGETEVENTS_TIME64         = 4416
	SYS_RECVMMSG_TIME64              = 4417
	SYS_MQ_TIMEDSEND_TIME64          = 4418
	SYS_MQ_TIMEDRECEIVE_TIME64       = 4419
	SYS_SEMTIMEDOP_TIME64            = 4420
	SYS_RT_SIGTIMEDWAIT_TIME64       = 4421
	SYS_FUTEX_TIME64                 = 4422
	SYS_SCHED_RR_GET_INTERVAL_TIME64 = 4423
	SYS_PIDFD_SEND_SIGNAL            = 4424
	SYS_IO_URING_SETUP               = 4425
	SYS_IO_URING_ENTER               = 4426
	SYS_IO_URING_REGISTER            = 4427
	SYS_OPEN_TREE                    = 4428
	SYS_MOVE_MOUNT                   = 4429
	SYS_FSOPEN                       = 4430
	SYS_FSCONFIG                     = 4431
	SYS_FSMOUNT                      = 4432
	SYS_FSPICK                       = 4433
	SYS_PIDFD_OPEN                   = 4434
	SYS_CLONE3                       = 4435
	SYS_CLOSE_RANGE                  = 4436
	SYS_OPENAT2                      = 4437
	SYS_PIDFD_GETFD                  = 4438
	SYS_FACCESSAT2                   = 4439
	SYS_PROCESS_MADVISE              = 4440
	SYS_EPOLL_PWAIT2                 = 4441
	SYS_MOUNT_SETATTR                = 4442
	SYS_QUOTACTL_FD                  = 4443
	SYS_LANDLOCK_CREATE_RULESET      = 4444
	SYS_LANDLOCK_ADD_RULE            = 4445
	SYS_LANDLOCK_RESTRICT_SELF       = 4446
	SYS_PROCESS_MRELEASE             = 4448
	SYS_FUTEX_WAITV                  = 4449
	SYS_SET_MEMPOLICY_HOME_NODE      = 4450
	SYS_CACHESTAT                    = 4451
	SYS_FCHMODAT2                    = 4452
	SYS_MAP_SHADOW_STACK             = 4453
	SYS_FUTEX_WAKE                   = 4454
	SYS_FUTEX_WAIT                   = 4455
	SYS_FUTEX_REQUEUE                = 4456
)

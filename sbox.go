package jh

var initVal224 = [16]uint64{
    0x2dfedd62f99a98ac, 0xae7cacd619d634e7,
    0xa4831005bc301216, 0xb86038c6c9661494,
    0x66d9899f2580706f, 0xce9ea31b1d9b1adc,
    0x11e8325f7b366e10, 0xf994857f02fa06c1,
    0x1b4f1b5cd8c840b3, 0x97f6a17f6e738099,
    0xdcdf93a5adeaa3d3, 0xa431e8dec9539a68,
    0x22b4a98aec86a1e4, 0xd574ac959ce56cf0,
    0x15960deab5ab2bbf, 0x9611dcf0dd64ea6e,
}

var initVal256 = [16]uint64{
    0xeb98a3412c20d3eb, 0x92cdbe7b9cb245c1,
    0x1c93519160d4c7fa, 0x260082d67e508a03,
    0xa4239e267726b945, 0xe0fb1a48d41a9477,
    0xcdb5ab26026b177a, 0x56f024420fff2fa8,
    0x71a396897f2e4d75, 0x1d144908f77de262,
    0x277695f776248f94, 0x87d5b6574780296c,
    0x5c5e272dac8e0d6c, 0x518450c657057a0f,
    0x7be4d367702412ea, 0x89e3ab13d31cd769,
}

var initVal384 = [16]uint64{
    0x481e3bc6d813398a, 0x6d3b5e894ade879b,
    0x63faea68d480ad2e, 0x332ccb21480f8267,
    0x98aec84d9082b928, 0xd455ea3041114249,
    0x36f555b2924847ec, 0xc7250a93baf43ce1,
    0x569b7f8a27db454c, 0x9efcbd496397af0e,
    0x589fc27d26aa80cd, 0x80c08b8c9deb2eda,
    0x8a7981e8f8d5373a, 0xf43967adddd17a71,
    0xa9b4d3bda475d394, 0x976c3fba9842737f,
}

var initVal512 = [16]uint64{
    0x6fd14b963e00aa17, 0x636a2e057a15d543,
    0x8a225e8d0c97ef0b, 0xe9341259f2b3c361,
    0x891da0c1536f801e, 0x2aa9056bea2b6d80,
    0x588eccdb2075baa6, 0xa90f3a76baf83bf7,
    0x0169e60541e34a69, 0x46b58a8e2e6fe65a,
    0x1047a7d0c1843c24, 0x3b6e71b12d5ac199,
    0xcf57f6ec9db1f856, 0xa706887c5716b156,
    0xe3c2fcdfe68517fb, 0x545a4678cc8cdd4b,
}

// =========

var C = []uint64{
    0x72d5dea2df15f867, 0x7b84150ab7231557,
    0x81abd6904d5a87f6, 0x4e9f4fc5c3d12b40,
    0xea983ae05c45fa9c, 0x03c5d29966b2999a,
    0x660296b4f2bb538a, 0xb556141a88dba231,
    0x03a35a5c9a190edb, 0x403fb20a87c14410,
    0x1c051980849e951d, 0x6f33ebad5ee7cddc,
    0x10ba139202bf6b41, 0xdc786515f7bb27d0,
    0x0a2c813937aa7850, 0x3f1abfd2410091d3,
    0x422d5a0df6cc7e90, 0xdd629f9c92c097ce,
    0x185ca70bc72b44ac, 0xd1df65d663c6fc23,
    0x976e6c039ee0b81a, 0x2105457e446ceca8,
    0xeef103bb5d8e61fa, 0xfd9697b294838197,
    0x4a8e8537db03302f, 0x2a678d2dfb9f6a95,
    0x8afe7381f8b8696c, 0x8ac77246c07f4214,
    0xc5f4158fbdc75ec4, 0x75446fa78f11bb80,
    0x52de75b7aee488bc, 0x82b8001e98a6a3f4,
    0x8ef48f33a9a36315, 0xaa5f5624d5b7f989,
    0xb6f1ed207c5ae0fd, 0x36cae95a06422c36,
    0xce2935434efe983d, 0x533af974739a4ba7,
    0xd0f51f596f4e8186, 0x0e9dad81afd85a9f,
    0xa7050667ee34626a, 0x8b0b28be6eb91727,
    0x47740726c680103f, 0xe0a07e6fc67e487b,
    0x0d550aa54af8a4c0, 0x91e3e79f978ef19e,
    0x8676728150608dd4, 0x7e9e5a41f3e5b062,
    0xfc9f1fec4054207a, 0xe3e41a00cef4c984,
    0x4fd794f59dfa95d8, 0x552e7e1124c354a5,
    0x5bdf7228bdfe6e28, 0x78f57fe20fa5c4b2,
    0x05897cefee49d32e, 0x447e9385eb28597f,
    0x705f6937b324314a, 0x5e8628f11dd6e465,
    0xc71b770451b920e7, 0x74fe43e823d4878a,
    0x7d29e8a3927694f2, 0xddcb7a099b30d9c1,
    0x1d1b30fb5bdc1be0, 0xda24494ff29c82bf,
    0xa4e7ba31b470bfff, 0x0d324405def8bc48,
    0x3baefc3253bbd339, 0x459fc3c1e0298ba0,
    0xe5c905fdf7ae090f, 0x947034124290f134,
    0xa271b701e344ed95, 0xe93b8e364f2f984a,
    0x88401d63a06cf615, 0x47c1444b8752afff,
    0x7ebb4af1e20ac630, 0x4670b6c5cc6e8ce6,
    0xa4d5a456bd4fca00, 0xda9d844bc83e18ae,
    0x7357ce453064d1ad, 0xe8a6ce68145c2567,
    0xa3da8cf2cb0ee116, 0x33e906589a94999a,
    0x1f60b220c26f847b, 0xd1ceac7fa0d18518,
    0x32595ba18ddd19d3, 0x509a1cc0aaa5b446,
    0x9f3d6367e4046bba, 0xf6ca19ab0b56ee7e,
    0x1fb179eaa9282174, 0xe9bdf7353b3651ee,
    0x1d57ac5a7550d376, 0x3a46c2fea37d7001,
    0xf735c1af98a4d842, 0x78edec209e6b6779,
    0x41836315ea3adba8, 0xfac33b4d32832c83,
    0xa7403b1f1c2747f3, 0x5940f034b72d769a,
    0xe73e4e6cd2214ffd, 0xb8fd8d39dc5759ef,
    0x8d9b0c492b49ebda, 0x5ba2d74968f3700d,
    0x7d3baed07a8d5584, 0xf5a5e9f0e4f88e65,
    0xa0b8a2f436103b53, 0x0ca8079e753eec5a,
    0x9168949256e8884f, 0x5bb05c55f8babc4c,
    0xe3bb3b99f387947b, 0x75daf4d6726b1c5d,
    0x64aeac28dc34b36d, 0x6c34a550b828db71,
    0xf861e2f2108d512a, 0xe3db643359dd75fc,
    0x1cacbcf143ce3fa2, 0x67bbd13c02e843b0,
    0x330a5bca8829a175, 0x7f34194db416535c,
    0x923b94c30e794d1e, 0x797475d7b6eeaf3f,
    0xeaa8d4f7be1a3921, 0x5cf47e094c232751,
    0x26a32453ba323cd2, 0x44a3174a6da6d5ad,
    0xb51d3ea6aff2c908, 0x83593d98916b3c56,
    0x4cf87ca17286604d, 0x46e23ecc086ec7f6,
    0x2f9833b3b1bc765e, 0x2bd666a5efc4e62a,
    0x06f4b6e8bec1d436, 0x74ee8215bcef2163,
    0xfdc14e0df453c969, 0xa77d5ac406585826,
    0x7ec1141606e0fa16, 0x7e90af3d28639d3f,
    0xd2c9f2e3009bd20c, 0x5faace30b7d40c30,
    0x742a5116f2e03298, 0x0deb30d8e3cef89a,
    0x4bc59e7bb5f17992, 0xff51e66e048668d3,
    0x9b234d57e6966731, 0xcce6a6f3170a7505,
    0xb17681d913326cce, 0x3c175284f805a262,
    0xf42bcbb378471547, 0xff46548223936a48,
    0x38df58074e5e6565, 0xf2fc7c89fc86508e,
    0x31702e44d00bca86, 0xf04009a23078474e,
    0x65a0ee39d1f73883, 0xf75ee937e42c3abd,
    0x2197b2260113f86f, 0xa344edd1ef9fdee7,
    0x8ba0df15762592d9, 0x3c85f7f612dc42be,
    0xd8a7ec7cab27b07e, 0x538d7ddaaa3ea8de,
    0xaa25ce93bd0269d8, 0x5af643fd1a7308f9,
    0xc05fefda174a19a5, 0x974d66334cfd216a,
    0x35b49831db411570, 0xea1e0fbbedcd549b,
    0x9ad063a151974072, 0xf6759dbf91476fe2,
};

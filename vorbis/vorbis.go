// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 01 May 2019 16:47:00 MSK.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package vorbis

/*
#cgo pkg-config: ogg vorbis vorbisenc
#include "ogg/ogg.h"
#include "vorbis/codec.h"
#include "vorbis/vorbisenc.h"
#include <stdlib.h>
#include <string.h>
#include <malloc.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// OggStreamPacketin function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_packetin.html
func OggStreamPacketin(os *OggStreamState, op *OggPacket) int32 {
	cos, _ := os.PassRef()
	cop, _ := op.PassRef()
	__ret := C.ogg_stream_packetin(cos, cop)
	__v := (int32)(__ret)
	return __v
}

func MyOggStreamPacketin(os *OggStreamState, op *OggPacket) int32 {

	var iop C.ogg_packet

	iop.packet = (*C.uchar)(C.calloc(C.size_t(op.Bytes), (C.size_t)(1)))

	C.memcpy(unsafe.Pointer(iop.packet), unsafe.Pointer(&op.Packet[0]), C.size_t(op.Bytes))

	iop.bytes = C.long(op.Bytes)
	iop.b_o_s = C.long(op.BOS)
	iop.e_o_s = C.long(op.EOS)

	C.memcpy(unsafe.Pointer(&iop.granulepos), unsafe.Pointer(&op.Granulepos), 8)
	C.memcpy(unsafe.Pointer(&iop.packetno), unsafe.Pointer(&op.Packetno), 8)

	cos, _ := os.PassRef()
	//cop, _ := op.PassRef()
	__ret := C.ogg_stream_packetin(cos, &iop)
	__v := (int32)(__ret)

	C.free(unsafe.Pointer(iop.packet))

	return __v
}

// OggStreamIovecin function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_iovecin.html
/*
func OggStreamIovecin(os *OggStreamState, iov *OggIovec, count int32, eOS int, granulepos OggInt64) int32 {
	cos, _ := os.PassRef()
	ciov, _ := iov.PassRef()
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	ceOS, _ := (C.long)(eOS), cgoAllocsUnknown
	cgranulepos, _ := (C.ogg_int64_t)(granulepos), cgoAllocsUnknown
	__ret := C.ogg_stream_iovecin(cos, ciov, ccount, ceOS, cgranulepos)
	__v := (int32)(__ret)
	return __v
}
*/

// OggStreamPageout function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_pageout.html
func OggStreamPageout(os *OggStreamState, og *OggPage) int32 {

	var iog OggPage

	cos, _ := os.PassRef()
	cog, _ := (&iog).PassRef()

	__ret := C.ogg_stream_pageout(cos, cog)

	if __ret > 0 {

		og.HeaderLen = int(iog.refb80411d1.header_len)
		og.Header = C.GoBytes(unsafe.Pointer(iog.refb80411d1.header), C.int(og.HeaderLen))
		og.BodyLen = int(iog.refb80411d1.body_len)
		og.Body = C.GoBytes(unsafe.Pointer(iog.refb80411d1.body), C.int(og.BodyLen))

	}

	__v := (int32)(__ret)
	return __v
}

// OggStreamPageoutFill function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_pageout_fill.html
func OggStreamPageoutFill(os *OggStreamState, og *OggPage, nfill int32) int32 {

	var iog OggPage

	cos, _ := os.PassRef()
	cog, _ := (&iog).PassRef()
	cnfill, _ := (C.int)(nfill), cgoAllocsUnknown
	__ret := C.ogg_stream_pageout_fill(cos, cog, cnfill)

	if __ret > 0 {
		og.HeaderLen = int(iog.refb80411d1.header_len)
		og.Header = C.GoBytes(unsafe.Pointer(iog.refb80411d1.header), C.int(og.HeaderLen))
		og.BodyLen = int(iog.refb80411d1.body_len)
		og.Body = C.GoBytes(unsafe.Pointer(iog.refb80411d1.body), C.int(og.BodyLen))
	}

	__v := (int32)(__ret)
	return __v
}

// OggStreamFlush function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_flush.html
func OggStreamFlush(os *OggStreamState, og *OggPage) int32 {

	var iog OggPage

	cos, _ := os.PassRef()
	cog, _ := (&iog).PassRef()
	__ret := C.ogg_stream_flush(cos, cog)

	if __ret > 0 {
		og.HeaderLen = int(iog.refb80411d1.header_len)
		og.Header = C.GoBytes(unsafe.Pointer(iog.refb80411d1.header), C.int(og.HeaderLen))
		og.BodyLen = int(iog.refb80411d1.body_len)
		og.Body = C.GoBytes(unsafe.Pointer(iog.refb80411d1.body), C.int(og.BodyLen))
	}

	__v := (int32)(__ret)
	return __v
}

// OggStreamFlushFill function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_flush_fill.html
func OggStreamFlushFill(os *OggStreamState, og *OggPage, nfill int32) int32 {
	cos, _ := os.PassRef()
	cog, _ := og.PassRef()
	cnfill, _ := (C.int)(nfill), cgoAllocsUnknown
	__ret := C.ogg_stream_flush_fill(cos, cog, cnfill)
	__v := (int32)(__ret)
	return __v
}

// OggSyncInit function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_init.html
func OggSyncInit(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_init(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncClear function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_clear.html
func OggSyncClear(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_clear(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncReset function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_reset.html
func OggSyncReset(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_reset(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncDestroy function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_destroy.html
func OggSyncDestroy(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_destroy(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncCheck function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_check.html
func OggSyncCheck(oy *OggSyncState) int32 {
	coy, _ := oy.PassRef()
	__ret := C.ogg_sync_check(coy)
	__v := (int32)(__ret)
	return __v
}

// OggSyncBuffer function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_buffer.html
func OggSyncBuffer(oy *OggSyncState, size int) []byte {
	coy, _ := oy.PassRef()
	csize, _ := (C.long)(size), cgoAllocsUnknown
	__ret := C.ogg_sync_buffer(coy, csize)
	__v := (*(*[0x7fffffff]byte)(unsafe.Pointer(__ret)))[:0]
	return __v
}

// OggSyncWrote function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_wrote.html
func OggSyncWrote(oy *OggSyncState, bytes int) int32 {
	coy, _ := oy.PassRef()
	cbytes, _ := (C.long)(bytes), cgoAllocsUnknown
	__ret := C.ogg_sync_wrote(coy, cbytes)
	__v := (int32)(__ret)
	return __v
}

// OggSyncPageseek function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_pageseek.html
func OggSyncPageseek(oy *OggSyncState, og *OggPage) int {
	coy, _ := oy.PassRef()
	cog, _ := og.PassRef()
	__ret := C.ogg_sync_pageseek(coy, cog)
	__v := (int)(__ret)
	return __v
}

// OggSyncPageout function as declared in https://xiph.org/ogg/doc/libogg/ogg_sync_pageout.html
func OggSyncPageout(oy *OggSyncState, og *OggPage) int32 {
	coy, _ := oy.PassRef()
	cog, _ := og.PassRef()
	__ret := C.ogg_sync_pageout(coy, cog)
	__v := (int32)(__ret)
	return __v
}

// OggStreamPagein function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_pagein.html
func OggStreamPagein(os *OggStreamState, og *OggPage) int32 {
	cos, _ := os.PassRef()
	cog, _ := og.PassRef()
	__ret := C.ogg_stream_pagein(cos, cog)
	__v := (int32)(__ret)
	return __v
}

// OggStreamPacketout function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_packetout.html
func OggStreamPacketout(os *OggStreamState, op *OggPacket) int32 {
	cos, _ := os.PassRef()
	cop, _ := op.PassRef()
	__ret := C.ogg_stream_packetout(cos, cop)
	__v := (int32)(__ret)
	return __v
}

// OggStreamPacketpeek function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_packetpeek.html
func OggStreamPacketpeek(os *OggStreamState, op *OggPacket) int32 {
	cos, _ := os.PassRef()
	cop, _ := op.PassRef()
	__ret := C.ogg_stream_packetpeek(cos, cop)
	__v := (int32)(__ret)
	return __v
}

// OggStreamInit function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_init.html
func OggStreamInit(os *OggStreamState, serialno int32) int32 {
	cos, _ := os.PassRef()
	cserialno, _ := (C.int)(serialno), cgoAllocsUnknown
	__ret := C.ogg_stream_init(cos, cserialno)
	__v := (int32)(__ret)
	return __v
}

// OggStreamClear function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_clear.html
func OggStreamClear(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_clear(cos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamReset function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_reset.html
func OggStreamReset(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_reset(cos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamResetSerialno function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_reset_serialno.html
func OggStreamResetSerialno(os *OggStreamState, serialno int32) int32 {
	cos, _ := os.PassRef()
	cserialno, _ := (C.int)(serialno), cgoAllocsUnknown
	__ret := C.ogg_stream_reset_serialno(cos, cserialno)
	__v := (int32)(__ret)
	return __v
}

// OggStreamDestroy function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_destroy.html
func OggStreamDestroy(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_destroy(cos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamCheck function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_check.html
func OggStreamCheck(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_check(cos)
	__v := (int32)(__ret)
	return __v
}

// OggStreamEos function as declared in https://xiph.org/ogg/doc/libogg/ogg_stream_eos.html
func OggStreamEos(os *OggStreamState) int32 {
	cos, _ := os.PassRef()
	__ret := C.ogg_stream_eos(cos)
	__v := (int32)(__ret)
	return __v
}

// OggPageChecksumSet function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_checksum_set.html
func OggPageChecksumSet(og *OggPage) {
	cog, _ := og.PassRef()
	C.ogg_page_checksum_set(cog)
}

// OggPageVersion function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_version.html
func OggPageVersion(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_version(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPageContinued function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_continued.html
func OggPageContinued(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_continued(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPageBos function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_bos.html
func OggPageBos(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_bos(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPageEos function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_eos.html
func OggPageEos(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_eos(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPageGranulepos function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_granulepos.html
func OggPageGranulepos(og *OggPage) OggInt64 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_granulepos(cog)
	__v := (OggInt64)(__ret)
	return __v
}

// OggPageSerialno function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_serialno.html
func OggPageSerialno(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_serialno(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPagePageno function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_pageno.html
func OggPagePageno(og *OggPage) int {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_pageno(cog)
	__v := (int)(__ret)
	return __v
}

// OggPagePackets function as declared in https://xiph.org/ogg/doc/libogg/ogg_page_packets.html
func OggPagePackets(og *OggPage) int32 {
	cog, _ := og.PassRef()
	__ret := C.ogg_page_packets(cog)
	__v := (int32)(__ret)
	return __v
}

// OggPacketClear function as declared in https://xiph.org/ogg/doc/libogg/ogg_packet_clear.html
func OggPacketClear(op *OggPacket) {
	cop, _ := op.PassRef()
	C.ogg_packet_clear(cop)
}

// InfoInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_info_init.html
func InfoInit(vi *Info) {
	cvi, _ := vi.PassRef()
	C.vorbis_info_init(cvi)
}

// InfoClear function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_info_clear.html
func InfoClear(vi *Info) {
	cvi, _ := vi.PassRef()
	C.vorbis_info_clear(cvi)
}

// InfoBlocksize function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_info_blocksize.html
func InfoBlocksize(vi *Info, zo int32) int32 {
	cvi, _ := vi.PassRef()
	czo, _ := (C.int)(zo), cgoAllocsUnknown
	__ret := C.vorbis_info_blocksize(cvi, czo)
	__v := (int32)(__ret)
	return __v
}

// CommentInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_init.html
func CommentInit(vc *Comment) {
	cvc, _ := vc.PassRef()
	C.vorbis_comment_init(cvc)
}

// CommentAdd function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_add.html
func CommentAdd(vc *Comment, comment string) {

	cmtlen := C.ulong(len(comment))

	cvc, _ := vc.PassRef()
	ccomment := (*C.char)(C.calloc(C.size_t(cmtlen+1), C.size_t(1)))

	scm := []byte(comment)
	scm = append(scm, 0)

	C.memcpy(unsafe.Pointer(ccomment), unsafe.Pointer(&scm[0]), C.size_t(cmtlen+1))
	C.vorbis_comment_add(cvc, ccomment)

	C.free(unsafe.Pointer(ccomment))
}

// CommentAddTag function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_add_tag.html
func CommentAddTag(vc *Comment, tag string, contents string) {
	cvc, _ := vc.PassRef()
	ctag, _ := unpackPCharString(tag)
	ccontents, _ := unpackPCharString(contents)
	C.vorbis_comment_add_tag(cvc, ctag, ccontents)
}

// CommentQuery function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_query.html
func CommentQuery(vc *Comment, tag string, count int32) *byte {
	cvc, _ := vc.PassRef()
	ctag, _ := unpackPCharString(tag)
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	__ret := C.vorbis_comment_query(cvc, ctag, ccount)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// CommentQueryCount function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_query_count.html
func CommentQueryCount(vc *Comment, tag string) int32 {
	cvc, _ := vc.PassRef()
	ctag, _ := unpackPCharString(tag)
	__ret := C.vorbis_comment_query_count(cvc, ctag)
	__v := (int32)(__ret)
	return __v
}

// CommentClear function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_comment_clear.html
func CommentClear(vc *Comment) {
	cvc, _ := vc.PassRef()
	C.vorbis_comment_clear(cvc)
}

// BlockInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_block_init.html
func BlockInit(v *DspState, vb *Block) int32 {
	cv, _ := v.PassRef()
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_block_init(cv, cvb)
	__v := (int32)(__ret)
	return __v
}

// BlockClear function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_block_clear.html
func BlockClear(vb *Block) int32 {
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_block_clear(cvb)
	__v := (int32)(__ret)
	return __v
}

// DspClear function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_dsp_clear.html
func DspClear(v *DspState) {
	cv, _ := v.PassRef()
	C.vorbis_dsp_clear(cv)
}

// GranuleTime function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_granule_time.html
func GranuleTime(v *DspState, granulepos OggInt64) float64 {
	cv, _ := v.PassRef()
	cgranulepos, _ := (C.ogg_int64_t)(granulepos), cgoAllocsUnknown
	__ret := C.vorbis_granule_time(cv, cgranulepos)
	__v := (float64)(__ret)
	return __v
}

// VersionString function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_version_string.html
func VersionString() string {
	__ret := C.vorbis_version_string()
	__v := packPCharString(__ret)
	return __v
}

// AnalysisInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_init.html
func AnalysisInit(v *DspState, vi *Info) int32 {
	cv, _ := v.PassRef()
	cvi, _ := vi.PassRef()
	__ret := C.vorbis_analysis_init(cv, cvi)
	__v := (int32)(__ret)
	return __v
}

// CommentheaderOut function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_commentheader_out.html
func CommentheaderOut(vc *Comment, op *OggPacket) int32 {
	var iop C.ogg_packet

	cvc, _ := vc.PassRef()
	__ret := C.vorbis_commentheader_out(cvc, &iop)

	op.Packet = C.GoBytes(unsafe.Pointer(iop.packet), C.int(iop.bytes))
	op.Bytes = int(iop.bytes)
	op.BOS = int(iop.b_o_s)
	op.EOS = int(iop.e_o_s)
	op.Granulepos = OggInt64(iop.granulepos)
	op.Packetno = OggInt64(iop.packetno)

	__v := (int32)(__ret)
	return __v
}

// AnalysisHeaderout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_headerout.html
func AnalysisHeaderout(v *DspState, vc *Comment, op *OggPacket, opComm *OggPacket, opCode *OggPacket) int32 {
	cv, _ := v.PassRef()
	cvc, _ := vc.PassRef()
	cop, _ := op.PassRef()
	copComm, _ := opComm.PassRef()
	copCode, _ := opCode.PassRef()
	__ret := C.vorbis_analysis_headerout(cv, cvc, cop, copComm, copCode)
	__v := (int32)(__ret)
	return __v
}

func PutUint32(b []byte, v uint32) []byte {
	var r []byte

	r = append(b, byte(v))
	r = append(r, byte(v>>8))
	r = append(r, byte(v>>16))
	r = append(r, byte(v>>24))

	return r
}

func PutUint16(b []byte, v uint16) []byte {
	var r []byte

	r = append(b, byte(v))
	r = append(r, byte(v>>8))

	return r
}

func PutInt16(b []byte, v int16) []byte {
	var r []byte
	r = append(b, byte(v))
	r = append(r, byte(v>>8))
	return r
}

// AnalysisHeaderout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_headerout.html
func OpusHeaderout(encoderName string, sample_rate int, channels int, op *OggPacket, opComm *OggPacket) int32 {

	var hdr OpusOGGHeader

	hdr.version = 1
	hdr.channel_count = channels
	hdr.pre_skip = 0
	hdr.input_sample_rate = uint32(sample_rate)
	hdr.output_gain = 0
	hdr.mapping_family = 0

	packetBuffer := []byte("OpusHead")
	packetBuffer = append(packetBuffer, byte(hdr.version))
	packetBuffer = append(packetBuffer, byte(hdr.channel_count))
	packetBuffer = PutUint16(packetBuffer, uint16(hdr.pre_skip))
	packetBuffer = PutUint32(packetBuffer, uint32(hdr.input_sample_rate))
	packetBuffer = PutInt16(packetBuffer, int16(hdr.output_gain))
	packetBuffer = append(packetBuffer, byte(hdr.mapping_family))

	if hdr.mapping_family != 0 {

		hdr.stream_count = 1
		hdr.coupled_count = 0
		hdr.mapping[0] = 0

		packetBuffer = append(packetBuffer, byte(hdr.stream_count))
		packetBuffer = append(packetBuffer, byte(hdr.coupled_count))
		for i := 0; i < channels; i++ {
			packetBuffer = append(packetBuffer, hdr.mapping[i])
		}
	}

	op.Packet = packetBuffer
	op.Bytes = len(packetBuffer)
	op.BOS = 1
	op.EOS = 0
	op.Granulepos = 0
	op.Packetno = 0

	packetBuffer = []byte("OpusTags")
	packetBuffer = PutUint32(packetBuffer, uint32(len(encoderName)))
	packetBuffer = append(packetBuffer, []byte(encoderName)...)
	packetBuffer = PutUint32(packetBuffer, uint32(0))

	opComm.Packet = packetBuffer
	opComm.Bytes = len(packetBuffer)
	opComm.BOS = 0
	opComm.EOS = 0
	opComm.Granulepos = 0
	opComm.Packetno = 1

	return 0
}

// AnalysisBuffer function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_buffer.html
func AnalysisBuffer(v *DspState, vals int32) **float32 {
	cv, _ := v.PassRef()
	cvals, _ := (C.int)(vals), cgoAllocsUnknown
	__ret := C.vorbis_analysis_buffer(cv, cvals)
	__v := *(***float32)(unsafe.Pointer(&__ret))
	return __v
}

// AnalysisWrote function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_wrote.html
func AnalysisWrote(v *DspState, vals int32) int32 {
	cv, _ := v.PassRef()
	cvals, _ := (C.int)(vals), cgoAllocsUnknown
	__ret := C.vorbis_analysis_wrote(cv, cvals)
	__v := (int32)(__ret)
	return __v
}

func AnalysisWriteBuffer(v *DspState, buffer []float32, vals int32) int32 {
	cv, _ := v.PassRef()
	cvals, _ := (C.int)(vals), cgoAllocsUnknown
	__ret := C.vorbis_analysis_buffer(cv, cvals)

	mychan := *__ret

	C.memcpy(unsafe.Pointer(mychan), unsafe.Pointer(&buffer[0]), C.size_t(vals)*4)

	___ret := C.vorbis_analysis_wrote(cv, cvals)
	__v := (int32)(___ret)
	return __v
}

// AnalysisBlockout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis_blockout.html
func AnalysisBlockout(v *DspState, vb *Block) int32 {
	cv, _ := v.PassRef()
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_analysis_blockout(cv, cvb)
	__v := (int32)(__ret)
	return __v
}

// Analysis function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_analysis.html
func Analysis(vb *Block, op *OggPacket) int32 {
	cvb, _ := vb.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_analysis(cvb, cop)
	__v := (int32)(__ret)
	return __v
}

// BitrateAddblock function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_bitrate_addblock.html
func BitrateAddblock(vb *Block) int32 {
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_bitrate_addblock(cvb)
	__v := (int32)(__ret)
	return __v
}

// BitrateFlushpacket function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_bitrate_flushpacket.html
func BitrateFlushpacket(vd *DspState, op *OggPacket) int32 {
	cvd, _ := vd.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_bitrate_flushpacket(cvd, cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisIdheader function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_idheader.html
func SynthesisIdheader(op *OggPacket) int32 {
	cop, _ := op.PassRef()
	__ret := C.vorbis_synthesis_idheader(cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisHeaderin function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_headerin.html
func SynthesisHeaderin(vi *Info, vc *Comment, op *OggPacket) int32 {
	cvi, _ := vi.PassRef()
	cvc, _ := vc.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_synthesis_headerin(cvi, cvc, cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_init.html
func SynthesisInit(v *DspState, vi *Info) int32 {
	cv, _ := v.PassRef()
	cvi, _ := vi.PassRef()
	__ret := C.vorbis_synthesis_init(cv, cvi)
	__v := (int32)(__ret)
	return __v
}

// SynthesisRestart function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_restart.html
func SynthesisRestart(v *DspState) int32 {
	cv, _ := v.PassRef()
	__ret := C.vorbis_synthesis_restart(cv)
	__v := (int32)(__ret)
	return __v
}

// Synthesis function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis.html
func Synthesis(vb *Block, op *OggPacket) int32 {
	cvb, _ := vb.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_synthesis(cvb, cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisTrackonly function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_trackonly.html
func SynthesisTrackonly(vb *Block, op *OggPacket) int32 {
	cvb, _ := vb.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_synthesis_trackonly(cvb, cop)
	__v := (int32)(__ret)
	return __v
}

// SynthesisBlockin function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_blockin.html
func SynthesisBlockin(v *DspState, vb *Block) int32 {
	cv, _ := v.PassRef()
	cvb, _ := vb.PassRef()
	__ret := C.vorbis_synthesis_blockin(cv, cvb)
	__v := (int32)(__ret)
	return __v
}

// SynthesisPcmout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_pcmout.html
func SynthesisPcmout(v *DspState, pcm [][][]float32) int32 {
	cv, _ := v.PassRef()
	cpcm, _ := unpackArgSSSFloat32(pcm)
	__ret := C.vorbis_synthesis_pcmout(cv, cpcm)
	packSSSFloat32(pcm, cpcm)
	__v := (int32)(__ret)
	return __v
}

// SynthesisLapout function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_lapout.html
func SynthesisLapout(v *DspState, pcm ***float32) int32 {
	cv, _ := v.PassRef()
	cpcm, _ := (***C.float)(unsafe.Pointer(pcm)), cgoAllocsUnknown
	__ret := C.vorbis_synthesis_lapout(cv, cpcm)
	__v := (int32)(__ret)
	return __v
}

// SynthesisRead function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_read.html
func SynthesisRead(v *DspState, samples int32) int32 {
	cv, _ := v.PassRef()
	csamples, _ := (C.int)(samples), cgoAllocsUnknown
	__ret := C.vorbis_synthesis_read(cv, csamples)
	__v := (int32)(__ret)
	return __v
}

// PacketBlocksize function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_packet_blocksize.html
func PacketBlocksize(vi *Info, op *OggPacket) int {
	cvi, _ := vi.PassRef()
	cop, _ := op.PassRef()
	__ret := C.vorbis_packet_blocksize(cvi, cop)
	__v := (int)(__ret)
	return __v
}

// SynthesisHalfrate function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_halfrate.html
func SynthesisHalfrate(v *Info, flag int32) int32 {
	cv, _ := v.PassRef()
	cflag, _ := (C.int)(flag), cgoAllocsUnknown
	__ret := C.vorbis_synthesis_halfrate(cv, cflag)
	__v := (int32)(__ret)
	return __v
}

// SynthesisHalfrateP function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_synthesis_halfrate_p.html
func SynthesisHalfrateP(v *Info) int32 {
	cv, _ := v.PassRef()
	__ret := C.vorbis_synthesis_halfrate_p(cv)
	__v := (int32)(__ret)
	return __v
}

// EncodeInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_encode_init.html
func EncodeInit(vi *Info, channels int, rate int, maxBitrate int, nominalBitrate int, minBitrate int) int32 {
	cvi, _ := vi.PassRef()
	cchannels, _ := (C.long)(channels), cgoAllocsUnknown
	crate, _ := (C.long)(rate), cgoAllocsUnknown
	cmaxBitrate, _ := (C.long)(maxBitrate), cgoAllocsUnknown
	cnominalBitrate, _ := (C.long)(nominalBitrate), cgoAllocsUnknown
	cminBitrate, _ := (C.long)(minBitrate), cgoAllocsUnknown
	__ret := C.vorbis_encode_init(cvi, cchannels, crate, cmaxBitrate, cnominalBitrate, cminBitrate)
	__v := (int32)(__ret)
	return __v
}

// EncodeSetupManaged function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_encode_setup_managed.html
func EncodeSetupManaged(vi *Info, channels int, rate int, maxBitrate int, nominalBitrate int, minBitrate int) int32 {
	cvi, _ := vi.PassRef()
	cchannels, _ := (C.long)(channels), cgoAllocsUnknown
	crate, _ := (C.long)(rate), cgoAllocsUnknown
	cmaxBitrate, _ := (C.long)(maxBitrate), cgoAllocsUnknown
	cnominalBitrate, _ := (C.long)(nominalBitrate), cgoAllocsUnknown
	cminBitrate, _ := (C.long)(minBitrate), cgoAllocsUnknown
	__ret := C.vorbis_encode_setup_managed(cvi, cchannels, crate, cmaxBitrate, cnominalBitrate, cminBitrate)
	__v := (int32)(__ret)
	return __v
}

// EncodeSetupVbr function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_encode_setup_vbr.html
func EncodeSetupVbr(vi *Info, channels int, rate int, quality float32) int32 {
	cvi, _ := vi.PassRef()
	cchannels, _ := (C.long)(channels), cgoAllocsUnknown
	crate, _ := (C.long)(rate), cgoAllocsUnknown
	cquality, _ := (C.float)(quality), cgoAllocsUnknown
	__ret := C.vorbis_encode_setup_vbr(cvi, cchannels, crate, cquality)
	__v := (int32)(__ret)
	return __v
}

// EncodeInitVbr function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_encode_init_vbr.html
func EncodeInitVbr(vi *Info, channels int, rate int, baseQuality float32) int32 {
	cvi, _ := vi.PassRef()
	cchannels, _ := (C.long)(channels), cgoAllocsUnknown
	crate, _ := (C.long)(rate), cgoAllocsUnknown
	cbaseQuality, _ := (C.float)(baseQuality), cgoAllocsUnknown
	__ret := C.vorbis_encode_init_vbr(cvi, cchannels, crate, cbaseQuality)
	__v := (int32)(__ret)
	return __v
}

// EncodeSetupInit function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_encode_setup_init.html
func EncodeSetupInit(vi *Info) int32 {
	cvi, _ := vi.PassRef()
	__ret := C.vorbis_encode_setup_init(cvi)
	__v := (int32)(__ret)
	return __v
}

// EncodeCtl function as declared in https://xiph.org/vorbis/doc/libvorbis/vorbis_encode_ctl.html
func EncodeCtl(vi *Info, number int32, arg unsafe.Pointer) int32 {
	cvi, _ := vi.PassRef()
	cnumber, _ := (C.int)(number), cgoAllocsUnknown
	carg, _ := arg, cgoAllocsUnknown
	__ret := C.vorbis_encode_ctl(cvi, cnumber, carg)
	__v := (int32)(__ret)
	return __v
}

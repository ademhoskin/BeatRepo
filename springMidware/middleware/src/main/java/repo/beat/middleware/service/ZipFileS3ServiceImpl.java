package repo.beat.middleware.service;

import org.springframework.http.*;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
public class ZipFileS3ServiceImpl implements ZipFileS3Service {

    @Autowired
    private RestTemplate restTemplate;

    @Override
    public ResponseEntity<String> uploadZipFile(String bucketName, String objectKey, String filePath) {
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.MULTIPART_FORM_DATA);

        MultiValueMap<String, Object> map = new LinkedMultiValueMap<>();
        map.add("bucketName", bucketName);
        map.add("objectKey", objectKey);
        map.add("filePath", filePath);
        HttpEntity<MultiValueMap<String, Object>> requestEntity = new HttpEntity<>(map, headers);

        return restTemplate.postForEntity("http://localhost:8080/s3/upload", requestEntity, String.class);

    }

    @Override
    public ResponseEntity<String> downloadZipFile(String bucketName, String objectKey, String filePath) {
        String url = "http://localhost:8080/s3/download?bucketName=" + bucketName + "&objectKey=" + objectKey + "&filePath=" + filePath;
        return restTemplate.getForEntity(url, String.class);
    }

    @Override
    public ResponseEntity<String> deleteZipFile(String bucketName, String keyName) {
        String url = "http://localhost:8080/s3/delete?bucketName=" + bucketName + "&keyName=" + keyName;
        return restTemplate.exchange(url, HttpMethod.DELETE, null, String.class);
    }
}
